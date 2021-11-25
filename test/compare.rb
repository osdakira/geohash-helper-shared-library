# fronzen_string_literal: true

require 'set'
require 'benchmark'

$LOAD_PATH.unshift(File.expand_path(File.join(File.dirname(__FILE__), "../lib")))
require 'geohash_helper'

def pure_ruby_intersect_geohashes(geohashes_a, geohashes_b)
  geohashes_and_sizes_a = geohashes_a.map { |x| [x, x.size - 1] }.to_h
  geohashes_and_sizes_b = geohashes_b.map { |x| [x, x.size - 1] }.to_h

  intersected = Set.new

  geohashes_and_sizes_b.each do |geohash_b, size_b|
    geohashes_and_sizes_a.each do |geohash_a, size_a|
      if size_a > size_b # a の方が長い
        if geohash_b == geohash_a[0..size_b] # a を b の桁に合わせたら一致した
          intersected.add(geohash_a)
        end
      elsif geohash_a == geohash_b[0..size_a] # b の方が長いときに、b を a の桁に合わせたら一致した
        intersected.add(geohash_b)
      end
    end
  end

  intersected.to_a
end

def make_geohash
  "0123456789bcdefghjkmnpqrstuvwxyz".split("").sample([*5..8].sample).join
end

geohashes_a = 20_000.times.map { make_geohash }
geohashes_b = 20_000.times.map { make_geohash }

Benchmark.bm do |r|
  r.report('ffi ') { GeohashHelper.intersect_geohashes(geohashes_a, geohashes_b) }
  r.report('ruby') { pure_ruby_intersect_geohashes(geohashes_a, geohashes_b) }
end
