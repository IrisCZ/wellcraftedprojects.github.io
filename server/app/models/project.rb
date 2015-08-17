class Project < ActiveRecord::Base
  validates :name, presence: true
  validates :url, presence: true
  validates :positives, numericality: {greater_than_or_equal_to: 0}
  validates :negatives, numericality: {greater_than_or_equal_to: 0}
  validates :image, format: {
    with: %r{\.(gif|jpg|png)\Z}i,
    message: 'must be a URL for GIF, JPG or PNG image'
  }

  after_initialize :set_defaults, unless: :persisted?

  def set_defaults
  	self.positives ||= 0
  	self.negatives ||= 0
  end
end
