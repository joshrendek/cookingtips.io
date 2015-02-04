root = exports ? this
{h1, button, label, ul, form, li, textarea, div, input} = React.DOM
NewPage = React.createClass
  displayName: 'NewPage'
  mixins: [window.pageForm]

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
      @showForm()

root.newPage = NewPage
