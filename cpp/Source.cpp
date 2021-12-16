#include <iostream>
#include <thread>
#include <queue>
#include <boost/lockfree/queue.hpp>

using namespace std;

const int loop = 1000;
queue<int> que;

void producer()
{
    for (int i = 0; i < loop; ++i)
    {
        que.push(i);
    }
}

void consumer()
{
}

int main()
{
    std::thread t1(producer);
    std::thread t2(producer);
    std::thread t3(producer);
    std::thread t4(producer);

    t1.join();
    t2.join();
    t3.join();
    t4.join();
    cout << que.size() << endl;
}
