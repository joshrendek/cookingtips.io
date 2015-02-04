root = exports ? this

{h1, button, label, ul, form, li, textarea, div, input} = React.DOM
PageForm =
  displayName: 'PageForm'

  showForm: ->
    div className: 'row',
      div className: 'form-group col-md-12',
        label {}, 'Title'
        input className: 'form-control', placeholder: 'Title', ref: 'title', value: @state?.page?.Title
      div className: 'form-group col-md-6',
        label {}, 'Instructions'
        textarea className: 'form-control', placeholder: 'Instructions, one per line', ref: 'instructions', value: @state?.page?.Instructions.join("\n")
      div className: 'form-group col-md-6',
        label {}, 'Youtube Videos'
        textarea className: 'form-control', placeholder: 'Youtube links, one per line', ref: 'youtubes', value: @state?.page?.Youtubes.join("\n")
      div className: 'form-group col-md-6',
        label {}, 'Articles'
        textarea className: 'form-control', placeholder: 'Articles, one per line', ref: 'articles', value: @state?.page?.Articles.join("\n")
      div className: 'form-group col-md-6',
        label {}, 'Tags'
        textarea className: 'form-control', placeholder: 'Tags, one per line', ref: 'tags', value: @state?.page?.Tags.join("\n")
      div className: 'form-group col-md-12',
        button className: 'btn btn-primary col-md-12', if @state?.page? then 'Save Page' else 'Add Page'

root.pageForm = PageForm
