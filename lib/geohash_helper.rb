# fronzen_string_literal: true

require 'fiddle/import'

require "geohash_helper/version"

module GeohashHelper
  extend Fiddle::Importer

  os = %x(uname -s).chomp.downcase
  arch = %x(uname -m).chomp.downcase
  arch = "amd64" if arch == "x86_64"
  dlload File.expand_path("./build/#{os}_#{arch}_geohash_lib_go.so", __dir__)

  extern 'int IsIntersect(char*, char*)'
  extern 'char* IntersectGeohashes(char** geohashes_a, int size_a, char** geohashes_b, int size_b)'
  extern 'void Free(void*)'

  def self.intersect?(geohash_a, geohash_b)
    IsIntersect(geohash_a, geohash_b) == 1 # 0,1 で返ってくる
  end

  def self.intersect_geohashes(geohashes_a, geohashes_b)
    # GC に回収されないように
    geohashes_a_pack = geohashes_a.pack('p*')
    geohashes_b_pack = geohashes_b.pack('p*')
    c_pointer = IntersectGeohashes(
      geohashes_a_pack, geohashes_a.size,
      geohashes_b_pack, geohashes_b.size
    )
    result_string = c_pointer.to_s
    Free(c_pointer)
    result_string.split(',') # csv 形式で配列が返ってくる
  end
end
