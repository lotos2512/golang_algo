package dijkstra

import (
	"container/heap"
)

type Price uint64

type Node struct {
	Name      string
	Relations []Relation
}

type TableData struct {
	node  *Node
	price Price
	from  *Node
	end   bool
}

type PriceTable map[string]*TableData

type Relation struct {
	Node  *Node
	Price Price
}

type Result struct {
	Nodes []*Node
	Price Price
}

func (n *Node) SetRelation(relation Relation) {
	n.Relations = append(n.Relations, relation)
}

func SearchBestWay(start, end *Node) (result Result) {
	if start == end {
		return Result{
			Nodes: []*Node{start},
			Price: 0,
		}
	}

	priceTable := make(PriceTable)

	priceTable[start.Name] = &TableData{
		price: 0,
		from:  nil,
		node:  start,
	}

	searchBestWay(start, end, priceTable)

	endPoint := priceTable.GetData(end.Name)

	if endPoint != nil {
		return priceTable.getResult(end)
	}
	return Result{
		Price: 0,
	}
}

func (pt PriceTable) getResult(node *Node) (result Result) {
	endNodePriceData := pt[node.Name]
	result.Nodes = []*Node{endNodePriceData.node}
	result.Price = endNodePriceData.price
	for endNodePriceData.from != nil {
		next := endNodePriceData.from
		result.Nodes = append(result.Nodes, next)
		endNodePriceData = pt.GetData(next.Name)
	}
	return
}

func searchBestWay(start, end *Node, priceTable PriceTable) {
	for start != nil {
		minIndex := 0
		var startPrice *TableData
		startPrice = priceTable.GetData(start.Name)
		if start != end {
			for i := range start.Relations {
				// получение цены
				relationPriceData := priceTable.GetData(start.Relations[i].Node.Name)
				if relationPriceData == nil {
					startPrice = priceTable.GetData(start.Name)
					relationPriceData = &TableData{
						price: startPrice.price + start.Relations[i].Price,
						from:  start,
						node:  start.Relations[i].Node,
					}
					priceTable[start.Relations[i].Node.Name] = relationPriceData
				}

				if relationPriceData.end || relationPriceData.node == startPrice.from {
					if minIndex == i {
						minIndex++
					}
					continue
				}

				nodePrice := start.Relations[i].Price + startPrice.price
				// проверка старого значения узла
				if nodePrice <= relationPriceData.price {
					relationPriceData.from = start
					relationPriceData.price = nodePrice
				}

				if start.Relations[i].Price < start.Relations[minIndex].Price {
					minIndex = i
				}
			}
		}
		startPrice.end = true

		if minIndex > len(start.Relations)-1 || end == start.Relations[minIndex].Node {
			start = priceTable.getNotReadyNode(end)
		} else {
			start = start.Relations[minIndex].Node
		}
	}
}

func (pt PriceTable) GetData(nodeName string) (priceData *TableData) {
	v, ok := pt[nodeName]
	if !ok {
		return
	}
	return v
}

func (pt PriceTable) getNotReadyNode(end *Node) *Node {
	// Создаем приоритетную очередь (минимальную кучу) для узлов
	pq := make(PriorityQueue, 0)

	// Инициализируем очередь данными из таблицы цен
	for _, v := range pt {
		if !v.end {
			heap.Push(&pq, v)
		}
	}

	for pq.Len() > 0 {
		// Извлекаем узел с минимальной стоимостью из очереди
		minNodeData := heap.Pop(&pq).(*TableData)

		// Если узел не равен конечному узлу и его стоимость меньше бесконечности, возвращаем его
		if minNodeData.node != end && minNodeData.price < Price(1<<63-1) {
			// Установим узлу признак того, что он уже обработан
			minNodeData.end = true
			return minNodeData.node
		}
	}

	return nil // Если все узлы уже обработаны или цена всех узлов равна бесконечности
}

// Дополнительные структуры для реализации минимальной кучи
type PriorityQueue []*TableData

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].price < pq[j].price }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*TableData)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
