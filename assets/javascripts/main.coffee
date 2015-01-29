root = exports ? this
{div, input} = React.DOM
Search = React.createClass
  displayName: 'Search'

  componentDidMount: ->
    @refs.searchInput.getDOMNode().focus()

  render: ->
    div className: 'search--form',
      input placeholder: 'Dice...', className: 'form-control', ref: 'searchInput'

root.search = Search
