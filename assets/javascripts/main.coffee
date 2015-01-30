root = exports ? this
{div, form, input} = React.DOM
Search = React.createClass
  displayName: 'Search'

  componentDidMount: ->
    @refs.searchInput.getDOMNode().focus()

  search: (e) ->
    e.preventDefault()
    #@refs.searchInput.getDOMNode().className = 'form-control transition-left'

  render: ->
    div className: 'search--form',
      form onSubmit: @search,
        input placeholder: 'Dice...', className: 'form-control', ref: 'searchInput'

root.search = Search
