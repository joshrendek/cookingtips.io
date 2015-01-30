root = exports ? this

{ul, li, div, input} = React.DOM

AdminMain = React.createClass
  displayName: 'AdminMain'
  render: ->
    div {},
      'Admin home page'

root.adminMain = AdminMain
