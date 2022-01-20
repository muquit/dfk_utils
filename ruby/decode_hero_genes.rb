#!/usr/bin/env ruby
########################################################################
# Decode defikingdoms starit or visual genes 
# Adapted from https://github.com/0rtis/dfk python code
# muquit@muquit.com Jan-18-2022 
########################################################################

require 'optparse'
require 'json'

require 'pp'

class DecodeDFKHeroGenes
  def initialize
    $stdout.sync = true
    $stderr.sync = true
    @me = File.basename($0)
    @alphabet = '123456789abcdefghijkmnopqrstuvwx'
    @debug = false
  end

  def log(msg)
    t = Time.new()
    puts "#{t}: #{msg}"
  end

  def log_debug(msg)
    return if !@debug
    $stderr.puts "(debug): #{msg}"
  end

  def rarity
    {
      0 =>  "common",
      1 =>  "uncommon",
      2 =>  "rare",
      3 =>  "legendary",
      4 =>  "mythic",
    }
  end

  def hclass
    {
      0 =>  "warrior",
      1 =>  "knight",
      2 =>  "thief",
      3 =>  "archer",
      4 =>  "priest",
      5 =>  "wizard",
      6 =>  "monk",
      7 =>  "pirate",
      16 => "paladin",
      17 => "darkKnight",
      18 => "summoner",
      19 => "ninja",
      24 => "dragoon",
      25 => "sage",
      28 => "dreadKnight"
    }
  end

  def visual_traits 
    {
      0 =>  'gender',
      1 =>  'headAppendage',
      2 =>  'backAppendage',
      3 =>  'background',
      4 =>  'hairStyle',
      5 =>  'hairColor',
      6 =>  'visualUnknown1',
      7 =>  'eyeColor',
      8 =>  'skinColor',
      9 =>  'appendageColor',
      10 => 'backAppendageColor',
      11 => 'visualUnknown2'
    }
  end

  def stat_traits
    {
      0 =>  'class',
      1 =>  'subClass',
      2 =>  'profession',
      3 =>  'passive1',
      4 =>  'passive2',
      5 =>  'active1',
      6 =>  'active2',
      7 =>  'statBoost1',
      8 =>  'statBoost2',
      9 =>  'statsUnknown1',
      10 => 'element',
      11 => 'statsUnknown2'
    }
  end

  def professions
    {
      0 =>  'mining',
      2 =>  'gardening',
      4 =>  'fishing',
      6 =>  'foraging',
    }
  end

  def stats 
    {
      0 =>  'strength',
      2 =>  'agility',
      4 =>  'intelligence',
      6 =>  'wisdom',
      8 =>  'luck',
      10 =>  'vitality',
      12 =>  'endurance',
      14 =>  'dexterity'
    }
  end

  def elements
    {
      0 =>  'fire',
      2 =>  'water',
      4 =>  'earth',
      6 =>  'wind',
      8 =>  'lightning',
      10 =>  'ice',
      12 =>  'light',
      14 =>  'dark',
    }
  end


  def alphabetize_genes(genes)
    base = @alphabet.length
    buf = ''
    while genes >= base
      mod = (genes % base)
      log_debug "#{genes} % #{base}=#{mod} #{genes.to_s.length}"
      buf = @alphabet[mod] + buf
      genes = (genes - mod) / base
      log_debug "genes: #{genes}"
    end
    buf = @alphabet[genes] + buf
    buf = buf.rjust(48, '1')
    log_debug "alphabetized genes: #{buf}"
    log_debug "Length: #{buf.length}"
    x = buf.scan(/.{1,4}/).join(" ")
    log_debug "#{x}"
    return buf
  end

  def decode_stat_genes(genes)
    stat_genes = {}
    stat_genes['type'] = 'Stat Genes'
    stat_genes['raw'] = genes
    abuf = alphabetize_genes(genes)
    log_debug "#{@alphabet}"
    i = 0
    abuf.split('').each_with_index do |val, idx|
      i = i + 1
      x = idx / 4
      stat_trait = stat_traits[x]
      k = abuf[idx]
      value_num = @alphabet.index(k)
      stat_genes[stat_trait] = value_num
      if i % 4 == 0 && @debug
        log_debug "idx: #{val} idx:#{idx} #{idx}/4:[#{x}] #{k} [#{value_num}] *"
        $stderr.puts "\n" 
      else
        log_debug "idx: #{val} #{idx} #{x} #{k} #{value_num}"
      end
    end
    stat_genes['class']         = hclass[stat_genes['class']]
    stat_genes['subClass']      = hclass[stat_genes['subClass']]
    stat_genes['profession']    = professions[stat_genes['profession']]
    stat_genes['statBoost1']    = stats[stat_genes['statBoost1']]
    stat_genes['statBoost2']    = stats[stat_genes['statBoost2']]
    stat_genes['statsUnknown1'] = stats[stat_genes['statsUnknown1']]
    stat_genes['statsUnknown2'] = stats[stat_genes['statsUnknown2']]
    stat_genes['element']       = elements[stat_genes['element']]
    return stat_genes
  end

  def decode_visual_genes(genes)
    visual_genes = {}
    visual_genes['type'] = 'Visual Genes'
    visual_genes['raw'] = genes
    abuf = alphabetize_genes(genes)
    abuf.split('').each_with_index do |val, idx|
      x = idx / 4
      stat_trait = visual_traits[x]
      k = abuf[idx]
      value_num = @alphabet.index(k)
      visual_genes[stat_trait] = value_num
    end
    visual_genes['gender'] = 'female'
    visual_genes['gender'] == 'male' if visual_genes['gender'] == 1
    return visual_genes
  end

  def doit
    options = {}
    ARGV << '-h' if ARGV.empty?
    OptionParser.new do |opts|
      opts.banner = "Usage: #{@me} [options]"
      opts.on("-s" " --stat-genes genes",
                "Decode stat genes") do |genes|
        options[:stat_genes] = genes
      end
      opts.on("-v" " --visual-genes genes",
                "Decode visual genes") do |genes|
        options[:visual_genes] = genes
      end
      opts.on("-d", "--debug", "Print debug messages") do 
        @debug = true
      end
      opts.on_tail("-h", "--help", "Show this message") do
        puts opts
        puts <<EOF
  Example:
  To decode Stat genes:
    #{$0} -s 55595053337262517174437940546058771473513094722680050621928661284030532
  To decode Visual genes:
    #{$0} -s 170877259497388213840353281232231169976585066888929467746175634464967719
EOF

        exit
      end
    end.parse!
    
#    if options[:stat_genes] == nil && options[:visual_genes] == nil
#      puts <<EOF
#  Usage: #{$0} -h for more info
#  Example:
#  To decode Stat genes:
#    #{$0} -s 55595053337262517174437940546058771473513094722680050621928661284030532
#  To decode Visual genes:
#    #{$0} -s 170877259497388213840353281232231169976585066888929467746175634464967719
#EOF
#      exit(1)
#    end

    if options[:stat_genes]
      genes = options[:stat_genes].to_i
      stat_genes = decode_stat_genes(genes)
      j = JSON.pretty_generate(stat_genes)
      puts(j)
    end

    if options[:visual_genes]
      genes = options[:visual_genes].to_i
      visual_genes = decode_visual_genes(genes)
      j = JSON.pretty_generate(visual_genes)
      puts(j)
    end
  end

end

if __FILE__ == $0
  DecodeDFKHeroGenes.new.doit()
end

