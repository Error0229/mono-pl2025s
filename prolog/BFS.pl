% Figure 11.10   An implementation of breadth-first search.


% solve( Start, Solution):
%    Solution is a path (in reverse order) from Start to a goal

solve( Start, Solution)  :-
  breadthfirst( [ [Start] ], Solution).

% breadthfirst( [ Path1, Path2, ...], Solution):
%   Solution is an extension to a goal of one of paths

breadthfirst( [ [Node | Path] | _], [Node | Path])  :-
  goal( Node).

breadthfirst( [Path | Paths], Solution)  :-
  extend( Path, NewPaths),
  conc( Paths, NewPaths, Paths1),
  breadthfirst( Paths1, Solution).

extend( [Node | Path], NewPaths)  :-
  bagof( [NewNode, Node | Path],
         ( s( Node, NewNode), not(member( NewNode, [Node | Path] )) ),
         NewPaths),
  !.

extend( Path, [] ).              % bagof failed: Node has no successor

conc([],L,L).
conc([X|L1],L2,[X|L3]) :-
  conc(L1,L2,L3).

s(a,b).
s(a,c).
s(b,d).
s(b,e).
s(c,f).
s(c,g).
s(d,h).
s(e,i).
s(e,j).
s(f,k).

goal(j).
goal(f).