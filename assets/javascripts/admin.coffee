root = exports ? this

{h1, button, label, ul, form, li, textarea, div, input} = React.DOM

NewPage = React.createClass
  displayName: 'NewPage'

  createPage: (e) ->
    e.preventDefault()
    console.log 'submitting'
    console.log @refs.title.getDOMNode().value
    $.ajax
      type: 'post'
      url: '/admin/pages/create'
      dataType: 'json'
      data:
        title: @refs.title.getDOMNode().value
        instructions: @refs.instructions.getDOMNode().value
        youtubes: @refs.youtubes.getDOMNode().value
        articles: @refs.articles.getDOMNode().value
        tags: @refs.tags.getDOMNode().value

      success: (e) ->
      done: (e) ->
        console.log e
      fail: (e) ->
        console.log e


  render: ->
    form onSubmit: @createPage, ref: 'newPageForm',
      h1 {}, 'Add New Page'
      div className: 'row',
        div className: 'form-group col-md-12',
          label {}, 'Title'
          input className: 'form-control', placeholder: 'Title', ref: 'title'
        div className: 'form-group col-md-6',
          label {}, 'Instructions'
          textarea className: 'form-control', placeholder: 'Instructions, one per line', ref: 'instructions'
        div className: 'form-group col-md-6',
          label {}, 'Youtube Videos'
          textarea className: 'form-control', placeholder: 'Youtube links, one per line', ref: 'youtubes'
        div className: 'form-group col-md-6',
          label {}, 'Articles'
          textarea className: 'form-control', placeholder: 'Articles, one per line', ref: 'articles'
        div className: 'form-group col-md-6',
          label {}, 'Tags'
          textarea className: 'form-control', placeholder: 'Tags, one per line', ref: 'tags'
        div className: 'form-group col-md-12',
          button className: 'btn btn-primary col-md-12', 'Add Page'


Admin = React.createClass
  displayName: 'Admin'

  getInitialState: ->
    displaying: AdminMain

  render: ->
    div className: 'container-fluid',
      div className: 'row admin--container',
        div className: 'col-md-2',
          ul {},
            li onClick: (=> @setState(displaying: NewPage)), 'Add Item',
        div className: 'col-md-10',
          @state.displaying()

root.admin = Admin
root.newPage = NewPage
