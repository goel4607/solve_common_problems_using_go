Let’s say we have a warehouse, that contains following things with expiry dates:
- Carrot that will expire in 30 days,
- Potato for 45 days,
- Rice for 100 Days and
- Spinach can be used in 5 days.

So to deliver these items, we must take care their expiry time, and to avoid the expiration, we will be delivering
these goods from the lowest to the highest expiry time.

So the order should be:

Item        Expiry in days
--------------------------
Carrot      30
Spinach     5
Rice        100
Potato      45

Expected:
--------------------------
Spinach     5
Carrot      30
Potato      45
Rice        100