ProjectListView = Backbone.View.extend({

  projects: new Projects(),

  initialize: function(){
    this.$el = $('#projects_container');
    this.listenTo(this.projects,'sync',this.render, this);
    this.listenTo(Backbone,'project:created', this.refresh, this);
    this.addProjectView = new AddProjectView({el:$('#modals')});
    this.refresh();
  },

  template: _.template($('#project_template').html()),

  render: function(){
    this.$el.html(this.template({projects:this.projects.toJSON()}));
    this.addProjectView.render();
  },

  refresh: function(){
    var collection = this.projects;
    this.projects.fetch({
      success: function(model, response) {
        collection.set(response.projects)
      }
    });
  }



});