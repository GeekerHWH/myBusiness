
# goods
| ID(gorm.Model) | Name | NewPrice | DiscountRate |
| --- | --- | --- | --- |
| 1 | Macbook Air M3 15'' | ¥10399 | 0.9023 |
| ... | ... | ... | ... |

> CREATE TABLE `goods` (
>     `id` integer PRIMARY KEY AUTOINCREMENT,
>     `created_at` datetime,
>     `updated_at` datetime,
>     `deleted_at` datetime,
>     `name` text NOT NULL,
>     `new_price` real NOT NULL,
>     `discount_rate` real NOT NULL
> );
>
> CREATE INDEX `idx_goods_deleted_at` ON `goods`(`deleted_at`);