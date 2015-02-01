(function() {
  var Admin, NewPage, button, div, form, h1, input, label, li, root, textarea, ul, _ref;

  root = typeof exports !== "undefined" && exports !== null ? exports : this;

  _ref = React.DOM, h1 = _ref.h1, button = _ref.button, label = _ref.label, ul = _ref.ul, form = _ref.form, li = _ref.li, textarea = _ref.textarea, div = _ref.div, input = _ref.input;

  NewPage = React.createClass({
    displayName: 'NewPage',
    createPage: function(e) {
      e.preventDefault();
      console.log('submitting');
      console.log(this.refs.title.getDOMNode().value);
      return $.ajax({
        type: 'post',
        url: '/admin/pages/create',
        dataType: 'json',
        data: {
          title: this.refs.title.getDOMNode().value,
          instructions: this.refs.instructions.getDOMNode().value,
          youtubes: this.refs.youtubes.getDOMNode().value,
          articles: this.refs.articles.getDOMNode().value,
          tags: this.refs.tags.getDOMNode().value
        },
        success: function(e) {},
        done: function(e) {
          return console.log(e);
        },
        fail: function(e) {
          return console.log(e);
        }
      });
    },
    render: function() {
      return form({
        onSubmit: this.createPage,
        ref: 'newPageForm'
      }, h1({}, 'Add New Page'), div({
        className: 'row'
      }, div({
        className: 'form-group col-md-12'
      }, label({}, 'Title'), input({
        className: 'form-control',
        placeholder: 'Title',
        ref: 'title'
      })), div({
        className: 'form-group col-md-6'
      }, label({}, 'Instructions'), textarea({
        className: 'form-control',
        placeholder: 'Instructions, one per line',
        ref: 'instructions'
      })), div({
        className: 'form-group col-md-6'
      }, label({}, 'Youtube Videos'), textarea({
        className: 'form-control',
        placeholder: 'Youtube links, one per line',
        ref: 'youtubes'
      })), div({
        className: 'form-group col-md-6'
      }, label({}, 'Articles'), textarea({
        className: 'form-control',
        placeholder: 'Articles, one per line',
        ref: 'articles'
      })), div({
        className: 'form-group col-md-6'
      }, label({}, 'Tags'), textarea({
        className: 'form-control',
        placeholder: 'Tags, one per line',
        ref: 'tags'
      })), div({
        className: 'form-group col-md-12'
      }, button({
        className: 'btn btn-primary col-md-12'
      }, 'Add Page'))));
    }
  });

  Admin = React.createClass({
    displayName: 'Admin',
    getInitialState: function() {
      return {
        displaying: AdminMain
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

  root.newPage = NewPage;

}).call(this);
