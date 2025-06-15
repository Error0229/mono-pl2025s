# 期末考題重點筆記

此筆記依據 `README.md` 最後段落列出的考試範圍撰寫，內容涵蓋 Prolog 狀態空間搜尋、啟發式搜尋技術，以及 Go 語言與經典論文中的併發概念。為方便短時間衝刺，每個主題都以說明加範例的方式呈現，總字數超過兩千字。

## 一、Prolog 與啟發式搜尋

### 1. 狀態空間表示
在 Prolog 中建模問題通常需定義狀態(space)與後繼關係(successor relation)。常見的寫法是 `s(X,Y,C)`，其中 `X`、`Y` 為狀態，`C` 為從 `X` 移至 `Y` 的成本。這種關係可透過事實或規則描述，並搭配搜尋演算法實現問題求解。為讓程式能自動探索可能狀態，我們也會定義目標狀態檢查，例如 `goal(State)`。

### 2. A* 演算法與評估函式
A* 演算法結合了成本累積函式 `g(n)` 與啟發式函式 `h(n)`，形成評估函式 `f(n) = g(n) + h(n)`。其中 `g(n)` 代表從起點到目前節點 `n` 的實際耗費，而 `h(n)` 則預估從 `n` 到目標的成本。若 `h(n)` 永遠不高於真實成本，則稱為「可容許」(admissible)，此時 A* 搜尋可以保證找到最短路徑。

A* 的流程包括：
1. 將起始節點放入開放表(open list)並計算 `f` 值。
2. 反覆從開放表取出 `f` 值最小的節點展開，將產生的新狀態根據 `f` 值插入表中。
3. 若開放表為空代表搜尋失敗，取出的節點若為目標則成功，否則持續。此過程會擴張大量節點，啟發式函式的好壞左右了搜尋效率。

### 3. 啟發式設計與可容許性
以八宮格(8-puzzle)為例，常見的啟發式包含：
- 曼哈頓距離：計算每個方塊離目標位置的水平距離與垂直距離之和。
- tile out of place：統計錯位方塊數量。
- 結合多項評估：例如將曼哈頓距離加上順序分數(sequence score)等。設計此類啟發式時須確保對任何節點都不高估真正成本，以維持可容許特性。

在實務上，我們可比較不同啟發式的節點展開數差異，體會「越精確的估計能越快導向目標」。此外，也要注意啟發式值的計算成本不可過高，否則得不償失。

### 4. IDA* 與其他變形
當搜尋空間龐大時，A* 需要儲存大量節點，容易導致記憶體爆炸。IDA*(Iterative Deepening A*) 透過逐漸提高 `f` 上限來控制展開深度，重複以深度優先方式搜尋，比傳統 A* 更省空間。另一種變形 RBFS (Recursive Best-First Search) 則在遞迴過程中記錄次佳 `f` 值，遇到更差情況即回溯，可在記憶體與時間之間取得折衷。

### 5. 排程問題的 A*
在多處理器排程問題中，狀態可以描述為每個處理器當前完成時間及剩餘任務集合。啟發式 `h(n)` 可能採用「所有剩餘任務平均分配到處理器後預估的最終完成時間」，進一步與目前最大完成時間 `Fin` 取差值：

```
Finall = (SUM(D_W) + SUM(F_J)) / m
h(n) = max(Finall - Fin, 0)
```

此啟發式在案例中被證明具可容許性，而且能大幅減少展開節點。若結合 branch and bound 或重複深化，更能抑制搜尋空間。

### 6. 範例：八宮格求解程式
以下是一個簡化的 Prolog A* 搜尋片段，示意如何結合 `s/3` 與 `h/2`：

```prolog
solve(Start, Path, Cost) :-
    astar([[Start,0]], [], Path, Cost).

astar([[State,G]|_], _, [State], G) :-
    goal(State).
astar([[State,G]|RestOpen], Closed, Path, Cost) :-
    findall([Next,G1], (s(State,Next,C), \+ member(Next,Closed), G1 is G+C), Children),
    append(RestOpen, Children, Open1),
    sort(2, @=<, Open1, OpenSorted),
    astar(OpenSorted, [State|Closed], PathRest, Cost),
    Path = [State|PathRest].
```

此程式每次取出 `f` 最小節點展開，並透過 `sort/4` 依第二欄位排序，雖然不夠完整但能說明基本概念。要具體應用，需要根據問題定義 `s/3` 與 `h/2`。

## 二、併發程式理論

### 1. 進程、同步與通訊
論文《Concepts and Notations for Concurrent Programming》總結了併發程式的核心概念：
- **進程(Process)**：獨立的執行單元，可視為「執行緒」或「任務」。
- **通訊(Communication)**：進程之間交換資料的方式，可透過共享變數或訊息傳遞。
- **同步(Synchronization)**：約束進程執行順序的手段，避免競爭或確保條件被滿足。

文章以操作式與公理式兩種角度探討同步需求。在操作式視角，程序是原子事件的序列；在公理式視角，可以用 Hoare Triple `{P} S {Q}` 推導程式正確性。

### 2. Busy-Waiting 與其缺點
早期的同步策略多利用忙等，即進程不斷輪詢共享變數以等待條件。這種方式雖易於理解，卻導致 CPU 空轉。Peterson 演算法雖能解決兩進程互斥問題，但仍需兩個旗標與一個 `turn` 變數，在實際系統中效率不佳。為此，後來出現 semaphore、monitor 等更高階機制。

### 3. Semaphores 與範例
Semaphore 提供 `P` (嘗試進入) 與 `V` (釋放) 兩操作。若 semaphore 值為 0，執行 `P` 的進程會被阻塞，直到其他進程 `V`。以下以伺服器處理輸入與輸出緩衝區的例子展示：

```pseudo
var in_mutex, out_mutex := semaphore(1)
var num_cards := semaphore(0)
var free_cards := semaphore(N)

process reader:
  loop
    read card
    P(free_cards); P(in_mutex)
    deposit card
    V(in_mutex); V(num_cards)
```

此例中的 `reader`、`executer`、`printer` 三個進程共用多個 semaphore，確保互斥與條件同步同時滿足。這類範式廣泛應用於作業系統核心或多線程程式庫。

### 4. 死鎖與公平
良好設計的同步機制需避免死鎖 (deadlock)。例如兩個進程同時持有對方所需的資源，若不釋放就會互相等待。Peterson 演算法透過 `turn` 變數確保不會永久等待，實現有限進度(finite progress)。一般而言，使用 semaphore 或 mutex 時，若取得鎖的順序不一致或忘記釋放鎖，都可能導致死鎖問題。

### 5. 訊息傳遞模式
論文也討論以訊息交換 (message passing) 實作同步，例如 CSP (Communicating Sequential Processes)。此模型強調「禁止共享變數」，改以通道傳送訊息。Go 語言的 goroutine 與 channel 即受到此思想啟發。訊息傳遞天然具有同步特性：發送者與接收者會在訊息交握時達成一致。

## 三、Go 語言併發特性

### 1. Goroutine 基本概念
Goroutine 是由 Go 執行時管理的輕量級執行緒。啟動 goroutine 時只需在函式呼叫前加 `go`，不必像傳統執行緒般預先配置棧大小。Go 執行時會根據需要動態調整棧容量，使成千上萬個 goroutine 在單機上仍可高效運行。

### 2. 通道(Channel) 與同步
Channel 是 goroutine 溝通的核心工具。創建通道常用 `make(chan T)`，亦可給定緩衝大小 `make(chan T, n)`。通道操作包含傳送 `ch <- v` 與接收 `v := <-ch`，在無緩衝情況下，兩者需同步完成。緩衝通道則允許有限度的異步傳送。範例 `pipeline1` 展示了三階段流水線：計數器、平方計算、輸出器，每一階段皆在獨立 goroutine 中進行。

### 3. 選擇(select) 與多路複用
`select` 陳述式允許 goroutine 同時監聽多個通道，根據就緒情況選擇執行。這對於處理多個事件來源或實作超時控制非常重要。例如火箭倒數程式使用 `time.Tick` 與 `abort` 兩通道，在 `select` 中依情況決定是否發射或取消。

```go
select {
case <-tick:
    fmt.Println(count)
case <-abort:
    fmt.Println("Launch aborted!")
    return
}
```

若多個 case 同時就緒，`select` 會隨機選擇其中之一，以避免某些通道長期被忽略。若欲實作「非阻塞」通訊，可搭配 `default` 分支。

### 4. WaitGroup 與工作同步
當需要等待多個 goroutine 結束時，可使用 `sync.WaitGroup`：

```go
var wg sync.WaitGroup
for _, f := range files {
    wg.Add(1)
    go func(name string) {
        defer wg.Done()
        thumbnail.ImageFile(name)
    }(f)
}
wg.Wait()
```

此例中 `wg.Wait()` 會在所有 goroutine 呼叫 `Done()` 後才返回，確保主程式不會太早結束。若 goroutine 在工作完畢前提前返回，也需透過 `defer wg.Done()` 保證計數正確遞減。

### 5. Mutex 與共享變數
若必須在多個 goroutine 間共享資料，可使用 `sync.Mutex`：

```go
var mu sync.Mutex
var balance int

func Deposit(amount int) {
    mu.Lock()
    balance += amount
    mu.Unlock()
}

func Balance() int {
    mu.Lock()
    b := balance
    mu.Unlock()
    return b
}
```

這類做法保護了共享變數 `balance`，但要注意：
1. 鎖定範圍應盡量小，避免阻塞過久。
2. 需確保任何錯誤路徑也會釋放鎖，通常用 `defer` 完成。

### 6. RWMutex 與效能考量
若讀取次數遠大於寫入，可改用 `sync.RWMutex`，允許多個讀取者同時進入臨界區，但寫入者仍需獨占：

```go
var mu sync.RWMutex

func ReadBalance() int {
    mu.RLock()
    defer mu.RUnlock()
    return balance
}

func WriteBalance(amount int) {
    mu.Lock()
    balance = amount
    mu.Unlock()
}
```

讀寫鎖在讀多寫少的情境下能提高吞吐量，但若寫入頻繁或鎖競爭低，則沒有明顯好處，甚至因維護額外狀態而慢些。

### 7. 記憶體模型與同步
現代 CPU 可能將寫入暫存於快取，導致不同核心觀察到的資料不一致。同步原語(包含 channel 操作與 mutex)不僅規定執行順序，也會刷新快取，確保其他處理器能看見最新資料。這是為何就算 `Balance` 僅讀取一個整數，也需要加鎖，以防止與 `Deposit` 的寫入交錯產生錯誤結果。

### 8. Goroutine 洩漏與取消
若 goroutine 因無人接收通道而永久阻塞，就會形成 goroutine 洩漏。可利用 `done` 通道或 `context.Context` 控制生命週期：

```go
done := make(chan struct{})

func worker() {
    for {
        select {
        case <-done:
            return
        default:
            // 執行任務
        }
    }
}

// 關閉 done 讓所有 worker 停止
close(done)
```

### 9. 綜合範例：併發爬蟲
結合前述技巧，可建立一個具限速與取消功能的並行網頁爬蟲：使用計數 semaphore 控制同時下載連結的數量，並在達到深度限制或收到取消信號時關閉通道，確保所有 goroutine 能順利退出。此例展示了 channels、goroutines、select 與 WaitGroup 的協同運用。

## 四、複習與建議

1. **動手實作八宮格或任務排程**：撰寫完整 A* 搜尋程式，並嘗試不同啟發式。實際運行後觀察節點展開差異，增進理解。
2. **閱讀論文與章節**：將《Concepts and Notations for Concurrent Programming》第一章及 3.1~3.2 詳讀並整理重點，尤其是 semaphore 範例與忙等缺點。
3. **撰寫小型併發程式**：如聊天室、檔案伺服器或計算密集工作，以熟悉 goroutine、channel、mutex 等工具。確保所有 goroutine 都會結束，以避免洩漏。
4. **自我測驗**：列出關鍵名詞（如 critical section、race condition、RWMutex 等），試著用自己的話解釋並寫下程式範例。
5. **維持規律**：每天安排固定時間複習並實作，搭配休息與自我檢驗。透過口頭講解或寫作，可更牢記概念。

## 結語

本筆記以超過兩千字的篇幅涵蓋了期末考指定內容：包括 Prolog 的狀態空間與 A* 搜尋、論文中的同步機制、Go 語言的併發模式，以及多種範例與實作要點。建議讀者在考前反覆練習程式和演算法推導，並確保掌握每個概念背後的理由與使用情境。熟悉這些主題後，相信能在期末考中順利發揮。

## 五、補充：Monitor 與訊息系統

除了 semaphore，論文也概述了 Conditional Critical Regions、Monitors 與 Path Expressions 等高階同步方法。Monitor 結合了資料與操作，只有透過 monitor 的程式碼才能存取內部資源，互斥與條件等待都由語言或執行環境隱含處理。Go 的 `sync.Cond` 與 `sync.Mutex` 可以模擬 monitor 風格，雖然語法上沒有顯式的 monitor 關鍵字，但透過封裝與方法呼叫即可達到同樣目的。

在訊息傳遞方面，論文介紹了同步、非同步通道，以及遠端程序呼叫(RPC)與交易 (Atomic Transaction) 觀念。Go 的 `net/rpc` 套件提供 RPC 機制，而更常用的 `net/http` 搭配 goroutine 也能實作分散式服務。若再加入 context 或自訂超時機制，便能處理複雜的網路錯誤情境。

## 六、進一步閱讀建議

1. **深入理解 A\***：可參考 AI 教科書，如 Russell & Norvig 的《Artificial Intelligence: A Modern Approach》，其中對 A\* 的分析更為詳盡，並討論一致性、一致啟發式等進階概念。
2. **CSP 與 Go**：欲瞭解 Go 通道背後的理論基礎，可研讀 Hoare 的《Communicating Sequential Processes》。雖然語法與 Go 不同，但核心思想相通。
3. **併發程式設計模式**：Go 社群有大量實作範例，如 worker pool、fan-in fan-out、pipeline 等。熟悉這些模式有助於實戰應用。

