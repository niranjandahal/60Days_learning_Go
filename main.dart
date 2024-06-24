class Solution {
  int networkDelayTime(List<List<int>> times, int n, int k) {
    // Create the adjacency list for the graph
    Map<int, List<List<int>>> graph = {};
    for (var time in times) {
      int u = time[0], v = time[1], w = time[2];
      if (!graph.containsKey(u)) {
        graph[u] = [];
      }
      graph[u]!.add([v, w]);
    }

    // Min-heap priority queue to store [time, node]
    PriorityQueue<List<int>> pq = PriorityQueue((a, b) => a[0] - b[0]);
    pq.add([0, k]); // Start with the source node k with time 0

    // Distances array to store the minimum time to reach each node
    List<int> dist = List.filled(n + 1, 1000000);
    dist[k] = 0;

    while (pq.isNotEmpty) {
      var current = pq.removeFirst();
      int currentTime = current[0];
      int node = current[1];

      if (currentTime > dist[node]) continue;

      if (graph.containsKey(node)) {
        for (var edge in graph[node]!) {
          int neighbor = edge[0];
          int time = edge[1];
          int newTime = currentTime + time;
          if (newTime < dist[neighbor]) {
            dist[neighbor] = newTime;
            pq.add([newTime, neighbor]);
          }
        }
      }
    }

    int maxTime = dist.skip(1).reduce((a, b) => a > b ? a : b);
    return maxTime == 1000000 ? -1 : maxTime;
  }
}

class PriorityQueue<T> {
  List<T> _heap;
  int Function(T a, T b) _compare;

  PriorityQueue(this._compare) : _heap = [];

  void add(T value) {
    _heap.add(value);
    _siftUp(_heap.length - 1);
  }

  T removeFirst() {
    if (_heap.isEmpty) {
      throw StateError('No element');
    }
    T first = _heap.first;
    if (_heap.length == 1) {
      _heap.removeLast();
    } else {
      _heap.first = _heap.removeLast();
      _siftDown(0);
    }
    return first;
  }

  bool get isEmpty => _heap.isEmpty;
  bool get isNotEmpty => _heap.isNotEmpty;

  void _siftUp(int index) {
    while (index > 0) {
      int parent = (index - 1) ~/ 2;
      if (_compare(_heap[index], _heap[parent]) >= 0) break;
      _swap(index, parent);
      index = parent;
    }
  }

  void _siftDown(int index) {
    int length = _heap.length;
    int leftChild = 2 * index + 1;
    while (leftChild < length) {
      int rightChild = leftChild + 1;
      int smallestChild = (rightChild < length &&
              _compare(_heap[rightChild], _heap[leftChild]) < 0)
          ? rightChild
          : leftChild;
      if (_compare(_heap[index], _heap[smallestChild]) <= 0) break;
      _swap(index, smallestChild);
      index = smallestChild;
      leftChild = 2 * index + 1;
    }
  }

  void _swap(int i, int j) {
    T temp = _heap[i];
    _heap[i] = _heap[j];
    _heap[j] = temp;
  }
}

void main() {
  Solution solution = Solution();

  // Example 1
  List<List<int>> times1 = [
    [2, 1, 1],
    [2, 3, 1],
    [3, 4, 1]
  ];
  print(solution.networkDelayTime(times1, 4, 2)); // Output: 2

  // Example 2
  List<List<int>> times2 = [
    [1, 2, 1]
  ];
  print(solution.networkDelayTime(times2, 2, 1)); // Output: 1

  // Example 3
  List<List<int>> times3 = [
    [1, 2, 1]
  ];
  print(solution.networkDelayTime(times3, 2, 2)); // Output: -1
}
