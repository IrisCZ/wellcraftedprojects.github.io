require 'test_helper'

class ProjectTest < ActiveSupport::TestCase

  test "name must not be empty"	do
    project = Project.new

    assert project.invalid?
    assert project.errors[:name].any?
    assert_equal [I18n.translate('errors.messages.blank')], project.errors[:name]
  end

  test "url must not be empty"	do
    project = Project.new

    assert project.invalid?
    assert project.errors[:url].any?
    assert_equal [I18n.translate('errors.messages.blank')], project.errors[:url]
  end

  test "posives attribute must be 0 by default" do
    project = Project.new

    assert_equal 0, project.positives
  end

  test "posives attribute should not changed to 0 if initialized it with any value" do
    project = Project.new({positives: 4})

    assert_equal 4, project.positives
  end

  test "posives attribute must be greater than or equal to 0" do
    project = Project.new

    project.positives = -1

    assert project.invalid?
    assert project.errors[:positives].any?
    assert_equal ["must be greater than or equal to 0"], project.errors[:positives]
  end

  test "negatives attribute must be 0 by default" do
    project = Project.new

    assert_equal 0, project.negatives
  end

  test "negatives attribute should not changed to 0 if initialized it with any value" do
    project = Project.new({negatives: 4})

    assert_equal 4, project.negatives
  end

  test "negatives attribute must be greater than or equal to 0" do
    project = Project.new

    project.negatives = -1

    assert project.invalid?
    assert project.errors[:negatives].any?
    assert_equal ["must be greater than or equal to 0"], project.errors[:negatives]
  end

  test "image should be a valid image" do
  	project = Project.new(name: "any", url: "http://any.com")
  	
  	assert project.invalid?
  	
  	project.image = "any.doc"
  	assert project.invalid?
    assert_equal ["must be a URL for GIF, JPG or PNG image"], project.errors[:image]

  	project.image = "folder/any.doc"
  	assert project.invalid?
    assert_equal ["must be a URL for GIF, JPG or PNG image"], project.errors[:image]

  	project.image = "any.jpg"
  	assert project.valid?

  	project.image = "ANY.gif"
  	assert project.valid?

  	project.image = "any.PNG"
  	assert project.valid?

  	project.image = "any/any.jpg"
  	assert project.valid?


  end

end
