defmodule Calibration do

  def get_data do
    System.cwd!
    |> Path.join("lib/data.txt")
    |> File.read
    |> case do
      {:ok, body}      -> String.split(body, "\n")
      {:error, reason} -> {:error, reason}
    end
  end

  def get_frequency(start \\ 0, list \\ []) do
    IO.puts("Starting Freq: #{start}")
    IO.puts("List.count: #{Enum.count(list)}")
    get_data
    |> Enum.reduce({start, list}, fn(offset_str, {acc, freq_list}) ->
      Integer.parse(offset_str)
      |> case do
        {offset, _} -> 
          new_freq = acc + offset
          cond do
            Enum.member?(freq_list, new_freq) ->
              raise "I Found your first duplicate! #{new_freq}"
            true -> { new_freq, [new_freq | freq_list] }
          end
        _ -> { acc, freq_list }
      end
    end)
  end

  def find_duplicate do
    find_duplicate({0, []})
  end
  def find_duplicate({start, list}) do
    {new_total, new_list} = get_frequency(start, list)
    find_duplicate({new_total, new_list})
  end

end
