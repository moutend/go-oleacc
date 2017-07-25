#!/usr/bin/env ruby

interface = ""
parent_interface = ""
methods = []

exit if ARGV.empty?

File.open(ARGV[0], "r").each_line do |line|
  tokens = line.chomp.split(" ")
  if tokens.empty?
    next
  end
  if interface == ""
    interface = tokens.first 
  end
  if parent_interface == ""
    parent_interface = tokens.last
  end
  if tokens.last.match(/\( *$/)
    m = tokens.last.gsub(/\(/, "").split("")
    m[0] = m[0].upcase
    methods << m.join("")
  end
end

puts <<heredoc
// #{interface.downcase}.go
package oleacc

import (
  "unsafe"

  "github.com/go-ole/go-ole"
)

type #{interface} struct {
  ole.#{parent_interface}
}

type #{interface}Vtbl struct {
  ole.#{parent_interface}Vtbl
heredoc

methods.each {|m| puts "  #{m} uintptr\n" }

puts <<heredoc
}

func (v *#{interface}) VTable() *#{interface}Vtbl{
  return (*#{interface}Vtbl)(unsafe.Pointer(v.RawVTable))
}

heredoc

methods.each do |method|
  puts <<heredoc
func (v *#{interface}) #{method}() error {
  return ole.NewError(ole.E_NOTIMPL)
}

heredoc
end
