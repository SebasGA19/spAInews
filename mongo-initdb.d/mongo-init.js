db.createUser(
    {
        user: "spainews",
        pwd: "spainews",
        roles: [
            {
                role: "readWrite",
                db: "spainews"
            }
        ]
    }
);