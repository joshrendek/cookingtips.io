root = exports ? this

{a, h1, button, label, ul, form, li, textarea, div, input} = React.DOM

ListPages = React.createClass
  displayName: 'ListPages'

  getInitialState: ->
    pages: []

  componentDidMount: ->
    $.ajax
      type: 'get'
      url: '/admin/pages'
      dataType: 'json'
      success: (data) =>
        @setState(pages: data)

  editPage: ->
    console.log 'edit'

  render: ->
    div {},
      @state.pages.map (page, index) =>
        div {},
          a onClick: @editPage, page.Title

Admin = React.createClass
  displayName: 'Admin'

  getInitialState: ->
    displaying: ListPages

  render: ->
    div className: 'container-fluid',
      div className: 'row admin--container',
        div className: 'col-md-2',
          ul {},
            li onClick: (=> @setState(displaying: NewPage)), 'Add Item',
        div className: 'col-md-10',
          @state.displaying()

root.admin = Admin
