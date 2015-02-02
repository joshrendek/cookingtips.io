(function() {
  var Admin, ListPages, a, button, div, form, h1, input, label, li, root, textarea, ul, _ref;

  root = typeof exports !== "undefined" && exports !== null ? exports : this;

  _ref = React.DOM, a = _ref.a, h1 = _ref.h1, button = _ref.button, label = _ref.label, ul = _ref.ul, form = _ref.form, li = _ref.li, textarea = _ref.textarea, div = _ref.div, input = _ref.input;

  ListPages = React.createClass({
    displayName: 'ListPages',
    getInitialState: function() {
      return {
        pages: []
      };
    },
    componentDidMount: function() {
      return $.ajax({
        type: 'get',
        url: '/admin/pages',
        dataType: 'json',
        success: (function(_this) {
          return function(data) {
            return _this.setState({
              pages: data
            });
          };
        })(this)
      });
    },
    editPage: function() {
      return console.log('edit');
    },
    render: function() {
      return div({}, this.state.pages.map((function(_this) {
        return function(page, index) {
          return div({}, a({
            onClick: _this.editPage
          }, page.Title));
        };
      })(this)));
    }
  });

  Admin = React.createClass({
    displayName: 'Admin',
    getInitialState: function() {
      return {
        displaying: ListPages
      };
    },
    render: function() {
      return div({
        className: 'container-fluid'
      }, div({
        className: 'row admin--container'
      }, div({
        className: 'col-md-2'
      }, ul({}, li({
        onClick: ((function(_this) {
          return function() {
            return _this.setState({
              displaying: NewPage
            });
          };
        })(this))
      }, 'Add Item'))), div({
        className: 'col-md-10'
      }, this.state.displaying())));
    }
  });

  root.admin = Admin;

}).call(this);
