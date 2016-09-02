class Log < ApplicationRecord
  enum flag: [info: 0, notice: 1, error: 2]
end
