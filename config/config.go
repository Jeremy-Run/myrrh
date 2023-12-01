package config

import "fmt"

var SimulationActivity = fmt.Sprintf(BaseActivity, SimulationExpr)

var BaseActivity = `
{
    "startTimestamp":1672502400000,
    "endTimestamp":1735660799000,
    "region":"SGP",
    "completionLimit":"OnceDuringActivity",
    "expr": %s,
    "action":"Ten percent off coupon"
}
`

var SimulationExpr = `
{
    "element":{

    },
    "logic":"||",
    "exprs":[
        {
            "element":{

            },
            "logic":"&&",
            "exprs":[
                {
                    "element":{
                        "feature":"LoginTimes",
                        "value":5,
                        "relation":">="
                    },
                    "logic":"",
                    "exprs":[

                    ]
                },
                {
                    "element":{
                        "feature":"OrderTimes",
                        "value":1,
                        "relation":">="
                    },
                    "logic":"",
                    "exprs":[

                    ]
                },
                {
                    "element":{
                        "feature":"CommentTimes",
                        "value":1,
                        "relation":">="
                    },
                    "logic":"",
                    "exprs":[

                    ]
                }
            ]
        },
        {
            "element":{

            },
            "logic":"&&",
            "exprs":[
                {
                    "element":{
                        "feature":"BrowseTimes",
                        "value":15,
                        "relation":">="
                    },
                    "logic":"",
                    "exprs":[

                    ]
                },
                {
                    "element":{
                        "feature":"PostTimes",
                        "value":3,
                        "relation":">="
                    },
                    "logic":"",
                    "exprs":[

                    ]
                },
                {
                    "element":{

                    },
                    "logic":"||",
                    "exprs":[
                        {
                            "element":{
                                "feature":"FollowTimes",
                                "value":8,
                                "relation":">="
                            },
                            "logic":"",
                            "exprs":[

                            ]
                        },
                        {
                            "element":{
                                "feature":"InvitationTimes",
                                "value":2,
                                "relation":">="
                            },
                            "logic":"",
                            "exprs":[

                            ]
                        }
                    ]
                }
            ]
        }
    ]
}
`
