defmodule CowInterrogator do
  def get_name do
    IO.gets("Waht is your name?")
    |> Strings.trim
  end

  def get_cow_lover do
    IO.getn("Do you like cows? [y|n] ", 1)
  end

  def interrogate do
    name = get_name()
    case String.downcase(get_cow_lover()) do
      "y" ->
        IO.puts "Great! Here's a cow for you #{name}:"
      "n" ->
        IO.puts "That's a shame, #{name}."
      _ ->
        IO.puts "You should have entered 'y' or 'n'."
    end
  end
end

CowInterrogator.interrogate
