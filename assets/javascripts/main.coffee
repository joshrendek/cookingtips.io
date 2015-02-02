root = exports ? this
{div, i, form, ul, li, input} = React.DOM
Search = React.createClass
  displayName: 'Search'

  componentDidMount: ->
    @refs.searchInput.getDOMNode().focus()

  getInitialState: ->
    results: null
    index: 0

  incrementIndex: ->
      pos = @state.index + 1
      pos = 0 if pos >= @state.results.length
      @setState(index: pos)

  goToItem: ->
    window.location = @state.results[@state.index].Url

  decrementIndex: ->
      pos = @state.index - 1
      pos = @state.results.length-1 if pos < 0
      @setState(index: pos)

  search: (e) ->
    @incrementIndex() if e.keyCode == 40
    @decrementIndex() if e.keyCode == 38
    @goToItem() if e.keyCode == 13
    q = @refs.searchInput.getDOMNode().value
    if q == ''
      @setState(results: null, index: 0)
      return
    e.preventDefault()
    $.ajax
      type: 'get'
      url: '/search'
      dataType: 'json'
      data:
        q: q
      success: (data) =>
        @setState(results: data)

    #@refs.searchInput.getDOMNode().className = 'form-control transition-left'

  clickItem: (index) ->
    window.location = @state.results[index].Url

  searchResults: ->
    div className: 'search--results',
      ul {},
        @state.results.map (page, index) =>
          className = if index == @state.index then 'search--selected' else ''
          li onClick: (=> @clickItem(index)),
            div className: "search--row #{className}",
              div className: 'search--title',
                page.Title
              div className: 'search--icons',
                div {}, "#{page.Instructions.length} Steps"
                div {}, "#{page.Youtubes.length} Videos"
                div {}, "#{page.Articles.length} Articles"

  render: ->
    div className: 'search--form',
      form onSubmit: @search,
        input placeholder: 'Dice...', className: 'form-control', ref: 'searchInput', onKeyUp: @search
      @searchResults() if @state.results?

root.search = Search
