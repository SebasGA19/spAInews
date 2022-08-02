db.createUser(
    {
        user: "api",
        pwd: "api",
        roles: [
            {
                role: "read",
                db: "spainews"
            }
        ]
    }
);