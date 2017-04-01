ExUnit.start

defmodule MyTest do
  use ExUnit.Case

  test "success test" do
    assert 1 + 1 == 2
  end

  test "failure test" do
    assert 1 + 1 == 3
  end

  test "refute is opposite of assert" do
    refute 1 + 1 == 3
  end

  test "assert raise" do
    assert_raise ArithmeticError, fn ->
      1 + "x"
    end
  end

  test "assert in delta" do
    assert_in_delta 1,
      5,
      6
  end
end

