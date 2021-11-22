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
      geohashes_a = %w[0 10 200 201 3001 3010 3100]

      expect(GeohashHelper.intersect_geohashes(geohashes_a, %w[1]).sort).to eq %w[10]
      expect(GeohashHelper.intersect_geohashes(geohashes_a, %w[2]).sort).to eq %w[200 201]
      expect(GeohashHelper.intersect_geohashes(geohashes_a, %w[20]).sort).to eq %w[200 201]
      expect(GeohashHelper.intersect_geohashes(geohashes_a, %w[202]).sort).to eq %w[]
      expect(GeohashHelper.intersect_geohashes(geohashes_a, %w[3]).sort).to eq %w[3001 3010 3100]
      expect(GeohashHelper.intersect_geohashes(geohashes_a, %w[4]).sort).to eq %w[]
      expect(GeohashHelper.intersect_geohashes(geohashes_a, %w[1 2]).sort).to eq %w[10 200 201]
    end

    context 'when the geohashes overlap by themselves' do
      it 'overlap is not removed' do
        # strictly expected = %w[101 102]
        geohashes_a = %w[10 101 102]
        expect(GeohashHelper.intersect_geohashes(geohashes_a, %w[10]).sort).to eq %w[10 101 102]

        # 012 has 01, but 01 has 0. 01 is not removed
        geohashes_a = %w[0 012]
        expect(GeohashHelper.intersect_geohashes(geohashes_a, %w[01]).sort).to eq %w[01 012]

        # overlap is removed
        geohashes_a = %w[012]
        expect(GeohashHelper.intersect_geohashes(geohashes_a, %w[01]).sort).to eq %w[012]
      end
    end
  end
end
