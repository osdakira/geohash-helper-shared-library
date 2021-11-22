# fronzen_string_literal: true

RSpec.describe GeohashHelper do
  describe ".intersect?" do
    it 'can check intersect for two geohashes' do
      expect(GeohashHelper.intersect?('012', '01')).to eq true
      expect(GeohashHelper.intersect?('012', '012')).to eq true
      expect(GeohashHelper.intersect?('012', '0123')).to eq true
      expect(GeohashHelper.intersect?('012', '01234')).to eq true
      expect(GeohashHelper.intersect?('012', '123')).to eq false
    end
  end

  describe ".intersect_geohashes" do
    it 'returns intersected geohashes' do
      geohashes_a = %w[0 02 0123 0124 012345]
      geohashes_b = %w[01 11]

      got = GeohashHelper.intersect_geohashes(geohashes_a, geohashes_b)
      expect(got.sort).to eq %w[01 0123 012345 0124]
    end
  end
end
