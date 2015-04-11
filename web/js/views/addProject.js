AddProjectView = Backbone.View.extend({

  model: new Project(),
  template: _.template($('#new_project_template').html()),
  container: $('#'),

  events: {
    'click #newproject-save':'submit'
  },

  initialize: function(){
  },

  render: function(){
    this.$el.html(this.template({}));
    return this;
  },

  submit: function(){
    this.model.set('name', this.$('#new-name').val().trim());
    this.model.set('url', this.$('#new-url').val().trim());
    this.model.set('author', this.$('#new-author').val().trim());
    this.model.set('tags', this.$('#new-tags').val().trim());
    this.model.set('description', this.$('#new-description').val().trim());
    this.model.save();
  }
});