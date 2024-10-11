package main

/*
각 점에 가중치가 부여된 트리가 주어집니다. 당신은 다음 연산을 통하여, 이 트리의 모든 점들의 가중치를 0으로 만들고자 합니다.

임의의 연결된 두 점을 골라서 한쪽은 1 증가시키고, 다른 한쪽은 1 감소시킵니다.
하지만, 모든 트리가 위의 행동을 통하여 모든 점들의 가중치를 0으로 만들 수 있는 것은 아닙니다. 당신은 주어진 트리에 대해서 해당 사항이 가능한지 판별하고, 만약 가능하다면 최소한의 행동을 통하여 모든 점들의 가중치를 0으로 만들고자 합니다.

트리의 각 점의 가중치를 의미하는 1차원 정수 배열 a와 트리의 간선 정보를 의미하는 edges가 매개변수로 주어집니다. 주어진 행동을 통해 트리의 모든 점들의 가중치를 0으로 만드는 것이 불가능하다면 -1을, 가능하다면 최소 몇 번만에 가능한지를 찾아 return 하도록 solution 함수를 완성해주세요. (만약 처음부터 트리의 모든 정점의 가중치가 0이라면, 0을 return 해야 합니다.)

제한사항
a의 길이는 2 이상 300,000 이하입니다.
a의 모든 수는 각각 -1,000,000 이상 1,000,000 이하입니다.
a[i]는 i번 정점의 가중치를 의미합니다.
edges의 행의 개수는 (a의 길이 - 1)입니다.
edges의 각 행은 [u, v] 2개의 정수로 이루어져 있으며, 이는 u번 정점과 v번 정점이 간선으로 연결되어 있음을 의미합니다.
edges가 나타내는 그래프는 항상 트리로 주어집니다.

[-5,0,2,1,2]	[[0,1],[3,4],[2,3],[0,3]]	9
*/

func solution(enroll []string, referral []string, seller []string, amount []int) []int {

	nodes := make(map[string]int)
	tree := make(map[int]int)

	rtn := make([]int, len(enroll))

	// enroll의 노드 번호 지정
	// root 노드 판별 및 자식 노드 식별
	for i := range enroll {
		nodes[enroll[i]] = i
		if referral[i] == "-" {
			tree[i] = -1
		} else {
			tree[i] = nodes[referral[i]] // 자식이 부모 노드 기억
		}
	}

	// 초기 수익 저장
	for i, s := range seller {
		dfs(nodes[s], 100*amount[i], rtn, tree)
	}
	// root 노드에서 시작해서 리프 노드로 이동. 리프 노드는 자신의 부모의 수익 갱신

	return rtn
}

func dfs(now int, gain int, profit []int, tree map[int]int) {

	off := gain / 10
	profit[now] += (gain - off)

	p := tree[now]
	if off != 0 && p != -1 {
		dfs(p, off, profit, tree)
	}
}
