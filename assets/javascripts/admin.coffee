root = exports ? this

{a, h1, button, label, ul, form, li, textarea, div, input} = React.DOM

ListPages = React.createClass
  displayName: 'ListPages'
  mixins: [window.pageForm]

  getInitialState: ->
    pages: []
    editing: null

  componentDidMount: ->
    $.ajax
      type: 'get'
      url: '/admin/pages'
      dataType: 'json'
      success: (data) =>
        @setState(pages: data, editing: null, page: null)

  editPage: (id) ->
    $.ajax
      type: 'get'
      url: '/admin/page/' + id
      dataType: 'json'
      success: (data) =>
        @setState(editing: id, page: data)

  render: ->
    div {},
      @state.pages.map (page, index) =>
        div {},
          a onClick: (=> @editPage(page.Id)), page.Title

Admin = React.createClass
  displayName: 'Admin'

  getInitialState: ->
    displaying: ListPages

  listPages: ->
    @setState(displaying: ListPages)

  render: ->
    div className: 'container-fluid',
      div className: 'row admin--container',
        div className: 'col-md-2',
          ul className: 'admin--nav',
            li {},
              a onClick: @listPages, 'List'
            li {},
              a onClick: (=> @setState(displaying: NewPage)), 'Add Item'
        div className: 'col-md-10',
          @state.displaying()


root.admin = Admin
root.listPages = ListPages
