# fronzen_string_literal: true

require 'test/unit'

$LOAD_PATH.unshift(File.expand_path(File.join(File.dirname(__FILE__), "../lib")))
require 'geohash_helper'

class GeohashHelperTest < Test::Unit::TestCase
  def test_intersect?
    assert_equal(GeohashHelper.intersect?('012', '01'), true)
    assert_equal(GeohashHelper.intersect?('012', '012'), true)
    assert_equal(GeohashHelper.intersect?('012', '0123'), true)
    assert_equal(GeohashHelper.intersect?('012', '01234'), true)
    assert_equal(GeohashHelper.intersect?('012', '123'), false)
  end

  def test_intersect_geohashes
    geohashes_a = %w[0 10 200 201 3001 3010 3100]

    assert_equal(GeohashHelper.intersect_geohashes(geohashes_a, %w[1]).sort, %w[10])
    assert_equal(GeohashHelper.intersect_geohashes(geohashes_a, %w[2]).sort, %w[200 201])
    assert_equal(GeohashHelper.intersect_geohashes(geohashes_a, %w[20]).sort, %w[200 201])
    assert_equal(GeohashHelper.intersect_geohashes(geohashes_a, %w[202]).sort, %w[])
    assert_equal(GeohashHelper.intersect_geohashes(geohashes_a, %w[3]).sort, %w[3001 3010 3100])
    assert_equal(GeohashHelper.intersect_geohashes(geohashes_a, %w[4]).sort, %w[])
    assert_equal(GeohashHelper.intersect_geohashes(geohashes_a, %w[1 2]).sort, %w[10 200 201])
  end

  def test_intersect_geohashes_when_the_geohashes_overlap_by_themselves
    test_cases = [
      { a: %w[10 101 102], b: %w[10], expected: %w[10 101 102] }, # strictly expected = %w[101 102]
      { a: %w[0 012], b: %w[01], expected: %w[01 012] }, # 012 has 01, but 01 has 0. 01 is not removed
      { a: %w[012], b: %w[01], expected: %w[012] }, # overlap is removed
    ]
    test_cases.each do |test_case|
      geohashes_a, geohashes_b, expected = test_case.values_at(:a, :b, :expected)
      assert_equal(GeohashHelper.intersect_geohashes(geohashes_a, geohashes_b).sort, expected)
    end
  end

  def test_make_gohashes_with_precision
    wkt_polygon = "POLYGON ((132.9709406620001 34.11831060600008, 132.9704386200001 34.11865367400003, 132.9697404750001 34.11975379000006, 132.96891048500004 34.11988086000002, 132.96853966600008 34.11865384600003, 132.96873584000002 34.11737278000004, 132.969587982 34.11632664000007, 132.97019897200005 34.116362571000025, 132.97015482900008 34.11731862700003, 132.9707228300001 34.11726456400004, 132.97113777200002 34.11758954100003, 132.9709406620001 34.11831060600008))"
    got = GeohashHelper.make_gohashes_with_precision(wkt_polygon, 7).sort
    expected = %w[wynd1f9 wynd1fb wynd1fc wynd1ff wynd1g0 wynd1g1 wynd1g2 wynd1g3 wynd1g4]
    assert_equal(got, expected)
  end
end
