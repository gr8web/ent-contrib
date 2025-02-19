// Copyright 2019-present Facebook
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package friendship

import (
	"time"

	"entgo.io/contrib/entgql/internal/todopulid/ent/predicate"
	"entgo.io/contrib/entgql/internal/todopulid/ent/schema/pulid"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldCreatedAt, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldUserID, v))
}

// FriendID applies equality check predicate on the "friend_id" field. It's identical to FriendIDEQ.
func FriendID(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldFriendID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Friendship {
	return predicate.Friendship(sql.FieldLTE(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v pulid.ID) predicate.Friendship {
	vc := string(v)
	return predicate.Friendship(sql.FieldContains(FieldUserID, vc))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v pulid.ID) predicate.Friendship {
	vc := string(v)
	return predicate.Friendship(sql.FieldHasPrefix(FieldUserID, vc))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v pulid.ID) predicate.Friendship {
	vc := string(v)
	return predicate.Friendship(sql.FieldHasSuffix(FieldUserID, vc))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v pulid.ID) predicate.Friendship {
	vc := string(v)
	return predicate.Friendship(sql.FieldEqualFold(FieldUserID, vc))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v pulid.ID) predicate.Friendship {
	vc := string(v)
	return predicate.Friendship(sql.FieldContainsFold(FieldUserID, vc))
}

// FriendIDEQ applies the EQ predicate on the "friend_id" field.
func FriendIDEQ(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldEQ(FieldFriendID, v))
}

// FriendIDNEQ applies the NEQ predicate on the "friend_id" field.
func FriendIDNEQ(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldNEQ(FieldFriendID, v))
}

// FriendIDIn applies the In predicate on the "friend_id" field.
func FriendIDIn(vs ...pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldIn(FieldFriendID, vs...))
}

// FriendIDNotIn applies the NotIn predicate on the "friend_id" field.
func FriendIDNotIn(vs ...pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldNotIn(FieldFriendID, vs...))
}

// FriendIDGT applies the GT predicate on the "friend_id" field.
func FriendIDGT(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldGT(FieldFriendID, v))
}

// FriendIDGTE applies the GTE predicate on the "friend_id" field.
func FriendIDGTE(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldGTE(FieldFriendID, v))
}

// FriendIDLT applies the LT predicate on the "friend_id" field.
func FriendIDLT(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldLT(FieldFriendID, v))
}

// FriendIDLTE applies the LTE predicate on the "friend_id" field.
func FriendIDLTE(v pulid.ID) predicate.Friendship {
	return predicate.Friendship(sql.FieldLTE(FieldFriendID, v))
}

// FriendIDContains applies the Contains predicate on the "friend_id" field.
func FriendIDContains(v pulid.ID) predicate.Friendship {
	vc := string(v)
	return predicate.Friendship(sql.FieldContains(FieldFriendID, vc))
}

// FriendIDHasPrefix applies the HasPrefix predicate on the "friend_id" field.
func FriendIDHasPrefix(v pulid.ID) predicate.Friendship {
	vc := string(v)
	return predicate.Friendship(sql.FieldHasPrefix(FieldFriendID, vc))
}

// FriendIDHasSuffix applies the HasSuffix predicate on the "friend_id" field.
func FriendIDHasSuffix(v pulid.ID) predicate.Friendship {
	vc := string(v)
	return predicate.Friendship(sql.FieldHasSuffix(FieldFriendID, vc))
}

// FriendIDEqualFold applies the EqualFold predicate on the "friend_id" field.
func FriendIDEqualFold(v pulid.ID) predicate.Friendship {
	vc := string(v)
	return predicate.Friendship(sql.FieldEqualFold(FieldFriendID, vc))
}

// FriendIDContainsFold applies the ContainsFold predicate on the "friend_id" field.
func FriendIDContainsFold(v pulid.ID) predicate.Friendship {
	vc := string(v)
	return predicate.Friendship(sql.FieldContainsFold(FieldFriendID, vc))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFriend applies the HasEdge predicate on the "friend" edge.
func HasFriend() predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FriendTable, FriendColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFriendWith applies the HasEdge predicate on the "friend" edge with a given conditions (other predicates).
func HasFriendWith(preds ...predicate.User) predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FriendInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, FriendTable, FriendColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Friendship) predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Friendship) predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Friendship) predicate.Friendship {
	return predicate.Friendship(func(s *sql.Selector) {
		p(s.Not())
	})
}
