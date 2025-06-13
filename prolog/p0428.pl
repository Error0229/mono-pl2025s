member1(X, [X|_]).
member1(X, [_|T]) :- member1(X, T).

conc([], L, L).
conc([H|T], L, [H|R]) :- conc(T, L, R).

member2(X, L) :- conc(_, [X|_], L).

parent( pam, bob). 
parent( tom, bob).
parent( tom, liz).
parent( bob, ann).
parent( bob, pat).
parent( pat, jim).         
 
predecessor( X, Z) :-   % Rule pr1
    parent( X, Z). 
 
predecessor( X, Z) :-   % Rule pr2
    parent( X, Y),
    predecessor( Y, Z).

predecessors(X, L) :- 
    findall(Y, predecessor(Y, X), L). 

% human(tom).

:- dynamic instance/2.

class(lane).
class(stage).
class(swimlane).

abstract(lane).
subclass(stage, lane).
subclass(swimlane, lane).

concrete(stage).
concrete(swimlane).

new_instance(I, C) :- 
    concrete(C),
    \+ instance(I, _),
    assert(instance(I, C)).