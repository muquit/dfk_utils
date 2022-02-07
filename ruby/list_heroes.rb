#!/usr/bin/env ruby
########################################################################
# List all of your heroes' attributes as JSON. The limit is 500 heroes
# at this time.
# Usage: list_heroes.rb <your address>
# muquit@muquit.com Dec-30-2021 
########################################################################
require 'net/http'
require 'json'
require 'uri'
require 'time'
require 'pp'

class ListHeroes
  def initialize
    $stdout.sync = true
    $stderr.sync = true
    @me = File.basename($0)
    @addr = nil
    n = 500
    @gql = """query {
      heros(first:#{n},orderBy: id, where: {owner: \"--addr--\"}) {
        owner {
          name
          id
          picId
        }
        id
        shiny
        shinyStyle
        statGenes
        visualGenes
        rarity
        firstName
        lastName
        mainClass
        subClass
        generation
        gardening
        mining
        fishing
        foraging
        strength
        intelligence
        dexterity
        endurance
        wisdom
        agility
        luck
        vitality
        mp
        hp
        stamina
        sp
        status
        staminaFullAt
        level
        xp
        currentQuest
        hpFullAt
        strengthGrowthP
        intelligenceGrowthP
        dexterityGrowthP
        enduranceGrowthP
        wisdomGrowthP
        agilityGrowthP
        luckGrowthP
        vitalityGrowthP
        hpSmGrowth
        hpRgGrowth
        hpLgGrowth
        mpSmGrowth
        mpRgGrowth
        mpLgGrowth
        strengthGrowthS
        intelligenceGrowthS
        dexterityGrowthS
        enduranceGrowthS
        wisdomGrowthS
        agilityGrowthS
        luckGrowthS
        vitalityGrowthS
        summonedTime
        maxSummons
        summons
        nextSummonTime
        summonerId {
          id
        }
        assistantId {
          id
        }
      }
    }"""
  end

  def log(msg)
    t = Time.new()
    puts "#{t}: #{msg}"
  end

  def get_json_object
    uri = URI("http://graph3.defikingdoms.com/subgraphs/name/defikingdoms/apiv5")
    @gql = @gql.gsub(/--addr--/,@addr)
    res = Net::HTTP.start(uri.host, uri.port) do |http|
      req = Net::HTTP::Post.new(uri)
      req['Content-Type'] = 'application/json'
      req.body = JSON[{'query' => @gql}]
      http.request(req)
    end
    json_str = res.body
    json_obj = JSON.parse(json_str)
    return json_obj
  end

  def check_args
    if ARGV.length != 1
      puts <<EOF
      Usage: #{@me} <your address>
EOF
      exit(1)
    end
    @addr = ARGV[0]
    # must be in lower case
    @addr = @addr.downcase
  end

  def print_stamina_full_at(json_obj)
    heroes = json_obj['data']['heros']
    heroes.each do |hero|
      stamina_full_at = hero['staminaFullAt']
      d = Time.at(stamina_full_at.to_i)
      t = Time.at(stamina_full_at.to_i).to_i
      ta = Time.at(stamina_full_at.to_i)
      t2 = Time.now().to_i
      diff =  t - t2
      id = hero['id']
      log "#{id}: Stamina full at: #{stamina_full_at}: #{ta} #{t}: #{diff}"
      log "#{id}: #{d}"
    end
  end

  def doit
    check_args()
    j = get_json_object()
    pp = JSON.pretty_generate(j)
    puts pp
    sz = j["data"]["heros"].size
    log "Number of heroes: #{sz}"
    print_stamina_full_at(j)
  end

end

if __FILE__ == $0
  ListHeroes.new.doit()
end

