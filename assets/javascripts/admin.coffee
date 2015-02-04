root = exports ? this

{a, h1, button, label, ul, form, li, textarea, div, input} = React.DOM

EditPage = React.createClass
  displayName: 'EditPage'
  mixins: [window.pageForm]

  getInitialState: ->
    state: null

  componentWillMount: ->
    $.ajax
      type: 'get'
      url: '/admin/page/' + @props.id
      dataType: 'json'
      success: (data) =>
        @setState(page: data)

  render: ->
    if @state.page? then @showForm() else null

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

  render: ->
    div {},
      @state.pages.map (page, index) =>
        div {},
          a href: "#/pages/#{page.Id}", page.Title

Admin = React.createClass
  displayName: 'Admin'

  render: ->
    div className: 'container-fluid',
      div className: 'row admin--container',
        div className: 'col-md-2',
          ul className: 'admin--nav',
            li {},
              a href: '#/pages', 'List'
            li {},
              a href: '#/pages/new', 'Add Page'
        div className: 'col-md-10',
          @props.displaying()


root.admin = Admin
root.listPages = ListPages
root.editPage = EditPage
