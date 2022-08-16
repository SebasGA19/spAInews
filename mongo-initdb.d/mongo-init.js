db.createUser(
    {
        user: "spainews",
        pwd: "spainews",
        roles: [
            {
                role: "read",
                db: "spainews"
            }
        ]
    }
);