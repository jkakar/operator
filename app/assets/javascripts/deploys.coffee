$ ->
  capitalize = (str) ->
    str[0].toLocaleUpperCase() + str.slice(1)

  buildOptionLabels = (property) ->
    if property.enum
      _.map property.enum, (v) -> v.split(":", 2)[0]
    else
      null

  buildFieldOptions = (properties) ->
    _.mapObject properties, (property, name) ->
      name: "options[#{name}]"
      title: property.title || capitalize(name)
      optionLabels: buildOptionLabels(property)

  $form = $("#deploy-form:first")
  $optionsForm = $("#deploy-options-form:first")
  optionFields = []
  if $optionsForm.length > 0
    validator = $optionsForm.data("validator")
    data = $optionsForm.data("data")
    data = null if _.isEmpty(data)

    $optionsForm.alpaca
      schema: validator
      data: data
      options:
        fields: buildFieldOptions(validator.properties)
      postRender: (renderedField) ->
        optionFields.push(renderedField)

  $form.on "submit", (e) ->
    for field in optionFields
      field.validate(true)
      field.refreshValidationState(true)
      unless field.isValid(true)
        e.preventDefault()
        return

    if confirm("Are you sure? This is your last chance to abort!")
      $form.find("input[type='submit']").disable()
    else
      e.preventDefault()
