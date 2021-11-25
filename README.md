# GeohashHelper

GeohashHelper is a gem for fast geohash handling.
It is a module that uses a shared library written by golang through fiddle.

## Installation

Add this line to your application's Gemfile:

```ruby
gem 'geohash_helper'
```

And then execute:

    $ bundle install

Or install it yourself as:

    $ gem install geohash_helper

## Usage

- `GeohashHelper.intersect?`

```rb
irb(main):008:0> GeohashHelper.intersect?('abc', 'ab')
=> true
irb(main):009:0> GeohashHelper.intersect?('abc', 'ac')
=> false
```

- `GeohashHelper.intersect_geohashes`

```
irb(main):010:0> GeohashHelper.intersect_geohashes(['abc', 'def'], ['abcdef', 'aa', 'de'])
=> ["def", "abcdef"]
```

## Speed

20,000 x 20,000 geohashes intersect comparison

```
$ ruby test/compare.rb
       user     system      total        real
ffi   7.350918   0.015016   7.365934 (  7.360118)
ruby 93.476806   0.143442  93.620248 ( 93.829111)
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/osdakira/geohash_helper.


## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).
