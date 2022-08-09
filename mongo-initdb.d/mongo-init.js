db.createUser(
    {
        user: "news-api",
        pwd: "news-api",
        roles: [
            {
                role: "read",
                db: "spainews"
            }
        ]
    }
);