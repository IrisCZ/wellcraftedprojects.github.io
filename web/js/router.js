Router = Backbone.Router.extend({
  routes: {
    '*actions': 'home'
  },

  initialize: function() {
    Backbone.history.start();
  },

  home: function(){
    new ProjectListView();
  }
});

var router = new Router();