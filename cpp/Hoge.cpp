#include <iostream>
#include <atomic>
#include <vector>
#include <thread>

int main()
{
  std::atomic<int> x(0);
  const int n = 300;

  auto lambda = [&x]
  {
    for (int i = 0; i < 3 * n; ++i)
    {
      int y = i;
      if (x == i)
      {
        x += 2;
        break;
      }
      /*
      if (x.compare_exchange_weak(y, y + 1) == true)
      {
        break;
      }
      */
    }
    return;
  };
  // いずれかのスレッドの処理がおわったら (成功したら) フラグをオンにする
  std::vector<std::thread> gogo;
  for (int i = 0; i < n; ++i)
  {
    gogo.push_back(std::thread(lambda));
  }
  for (int i = 0; i < n; ++i)
  {
    gogo[i].join();
  }

  std::cout << std::boolalpha << x << std::endl;
}
