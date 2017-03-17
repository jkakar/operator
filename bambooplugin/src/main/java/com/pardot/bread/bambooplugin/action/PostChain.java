package com.pardot.bread.bambooplugin.action;

import com.atlassian.bamboo.chains.Chain;
import com.atlassian.bamboo.chains.ChainExecution;
import com.atlassian.bamboo.chains.ChainResultsSummary;
import com.atlassian.bamboo.chains.StageExecution;
import com.atlassian.bamboo.chains.branches.MergeResultState;
import com.atlassian.bamboo.chains.branches.MergeResultSummary;
import com.atlassian.bamboo.chains.plugins.PostChainAction;
import com.atlassian.bamboo.configuration.AdministrationConfigurationAccessor;
import com.atlassian.bamboo.plan.PlanHelper;
import com.atlassian.bamboo.plan.PlanKey;
import com.atlassian.bamboo.plan.PlanResultKey;
import com.atlassian.bamboo.plan.cache.CachedPlanManager;
import com.atlassian.bamboo.plan.cache.ImmutablePlan;
import com.atlassian.bamboo.plugins.git.GitHubRepository;
import com.atlassian.bamboo.repository.Repository;
import com.atlassian.bamboo.repository.RepositoryData;
import com.atlassian.bamboo.security.EncryptionService;
import com.atlassian.bamboo.utils.BambooUrl;
import com.pardot.bread.bambooplugin.GithubStatusBuildConfiguration;
import com.pardot.bread.bambooplugin.GithubStatus;
import org.apache.log4j.Logger;
import org.jetbrains.annotations.NotNull;
import org.kohsuke.github.GHCommitState;

import java.io.IOException;

public class PostChain implements PostChainAction {
    private static Logger log = Logger.getLogger(PostChain.class);

    private AdministrationConfigurationAccessor administrationConfigurationAccessor;
    private CachedPlanManager cachedPlanManager;
    private EncryptionService encryptionService;

    public void setAdministrationConfigurationAccessor(AdministrationConfigurationAccessor administrationConfigurationAccessor) {
        this.administrationConfigurationAccessor = administrationConfigurationAccessor;
    }

    public void setCachedPlanManager(CachedPlanManager cachedPlanManager) {
        this.cachedPlanManager = cachedPlanManager;
    }

    public void setEncryptionService(EncryptionService encryptionService) {
        this.encryptionService = encryptionService;
    }

    @Override
    public void execute(@NotNull Chain chain, @NotNull ChainResultsSummary chainResultsSummary, @NotNull ChainExecution chainExecution) throws InterruptedException, Exception {
        setGithubStatusForMerge(chainResultsSummary, chainExecution);
    }

    private void setGithubStatusForMerge(final ChainResultsSummary chainResultsSummary, final ChainExecution chainExecution) {
        MergeResultSummary mergeResultSummary = chainResultsSummary.getMergeResult();
        if (mergeResultSummary == null || mergeResultSummary.getMergeState() != MergeResultState.SUCCESS) {
            return;
        }

        final String sha = mergeResultSummary.getMergeResultVcsKey();
        if (sha == null) {
            return;
        }

        // Backfill all of the statuses on the merge commit SHA
        for (final StageExecution stageExecution : chainExecution.getStages()) {
            if (stageExecution.isSuccessful()) {
                setGithubStatusForMerge(stageExecution, sha, GHCommitState.SUCCESS);
            } else {
                setGithubStatusForMerge(stageExecution, sha, GHCommitState.FAILURE);
            }
        }
    }

    private void setGithubStatusForMerge(final StageExecution stageExecution, final String sha, final GHCommitState state) {
        final ChainExecution chainExecution = stageExecution.getChainExecution();
        final PlanResultKey planResultKey = chainExecution.getPlanResultKey();
        final PlanKey planKey = planResultKey.getPlanKey();
        final ImmutablePlan plan = cachedPlanManager.getPlanByKey(planKey);

        if (plan == null) {
            return;
        }

        if (!GithubStatusBuildConfiguration.isGithubStatusEnabled(plan.getBuildDefinition().getCustomConfiguration())) {
            return;
        }

        final RepositoryData defaultRepositoryData = PlanHelper.getDefaultRepositoryDefinition(plan);
        if (defaultRepositoryData == null) {
            return;
        }

        final Repository defaultRepository = defaultRepositoryData.getRepository();
        if (defaultRepository instanceof GitHubRepository) {
            final GitHubRepository gitHubRepository = (GitHubRepository) defaultRepository;
            final String url = new BambooUrl(administrationConfigurationAccessor).withBaseUrlFromConfiguration("/browse/" + planResultKey.toString());

            try {
                log.info("Setting GitHub commit status for '" + planKey + "'/'" + stageExecution.getName() + "' to " + state);
                GithubStatus.create(
                        gitHubRepository,
                        encryptionService,
                        sha,
                        state,
                        url,
                        stageExecution.getName()
                );
            } catch (IOException e) {
                log.error("Unable to set GitHub commit status: " + e);
            }
        }
    }
}
