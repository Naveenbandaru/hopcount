# hopcount
Latency Aware Data Partitioning Techniques for Distributed Systems

### Paper Information
- **Author(s):** Naveen Kumar Bandaru
- **Published In:** International Journal of Innovative Research and Creative Technology (IJIRCT)
- **Publication Date:** June, 2023
- **E-ISSN:**  2454-5988

### Abstract
Data partitioning significantly influences communication efficiency in distributed systems by determining how requests access data across multiple nodes. Conventional static partitioning methods often ignore access locality, causing requests to traverse several intermediate nodes and increasing hop count. This study examines the impact of partition placement on communication distance and overall system efficiency. A latency aware partitioning approach is presented to route requests to nearby storage nodes and reduce unnecessary network traversal. Experimental analysis across different cluster sizes shows that the proposed approach significantly lowers hop count and improves scalability in distributed environments.

### Core Technical Contributions
- **Latency Conscious Partition Placement Approach:**  
Presented a partition placement method that incorporates communication locality and routing distance to minimize hop count during distributed data access operations.
- **Locality Based Intelligent Request Routing:**  
Developed a routing mechanism that dynamically directs requests toward the closest storage node by analyzing runtime communication paths and locality patterns.
- **Layered Distributed System Simulation Framework:** 
Constructed a distributed system prototype using Go based concurrent processing to simulate realistic request flows across multiple nodes and evaluate communication behavior.
- **Cluster Scale Communication Analysis:**  
Performed experiments on clusters consisting of 3, 5, 7, 9, and 11 nodes to study how hop distance and communication efficiency evolve as distributed environments expand.

### System Level Relevance and Benefits
- **Lower Network Traversal Distance:**
The latency aware placement strategy significantly decreases the number of intermediate nodes a request must traverse compared with static partition placement.

- **Enhanced Communication Performance:**  
Shorter routing paths reduce network delay and minimize processing overhead introduced by intermediate nodes in distributed clusters.

- **Improved Scalability for Large Systems:**  
The proposed approach maintains controlled hop growth even as cluster size increases, supporting stable performance in large scale distributed deployments.

- **Applicability to Modern Distributed Platforms:**  
The design can benefit distributed databases, cloud storage infrastructures, analytics systems, and microservice architectures that rely on efficient cross node data access.

### Experimental Results (Summary)

  | Nodes | Static Hops | Latency aware Hops | Improvment (%) |
  |-------|-------------| -------------------| ---------------|
  | 3     |  2.8        | 1.5                | 46.43          |
  | 5     |  3.4        | 1.7                | 50.00          |
  | 7     |  4.0        | 1.9                | 52.50          |
  | 9     |  4.6        | 2.1                | 54.35          |
  | 11    |  5.1        | 2.3                | 54.90          |

### Citation
Latency Aware Data Partitioning Techniques for Distributed Systems.
* Naveen Kumar Bandaru
* International Journal of Innovative Research and Creative Technology (IJIRCT) 
* E-ISSN 2454-5988
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijirct.org/ \
**Author Contact** \
**LinkedIn**: linkedin.com/in/naveen-bandaru | **Email**: naveen.bandaru@gmail.com







