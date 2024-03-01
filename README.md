## Architecture :-

API
1. Create User - Returns User ID
2. Login User and get JWT Token
3. CRUD Images and Playlist





[OPTIONAL]
In My domain ( digitalsignage.com/api/v1/configs ) will return the IP of the rabbit MQ and other server specific config option

Token Based Authentication (JWT Auth) is mandatory for communication

User Creation :-
Whenever a user is created, 
1. a new vhost ( multitenancy in rabbitmq ) will be ceated and 
2. a new playlist collection of that specific user will be created
3. a new bucket in object store will be created specific to the user


Image Store :-
Image will be stored in an Object Store ( minio or AWS S3) and the public link of the image will be used throughout the application

Block Type in Screen :-
lets say the whole screen is identified as block B0, in a split screen scenario one screen is identified as B1 and the other as B2 for three splits it can be B1 and B2TOP and B2BOTTOM

User to Device Mapping :-
Each client can have multiple devices connected. Each device will have a unique ID and each device can show different things

DataBases :-
Screen_Block_Type - It Stores block name ( B0,B1 etc ) and the aspect ratio of each block type .
Users - To store user information ( along with a unique user ID)
Images - A Database to store Image Metadata ( User Id, Image link, upload date and so on)
Playlists :- A collection of collection ( one collection for each user )
-   Each Collection stores Single Image or Image playlist that needs to be displayed for that particular User

Data Format to be stored in Playlist NoSQL database :-
For Storing Single Image ( Only one image will be displayed on the screen)
{
    id :- Unique_ID
    type :- Single
    device_id :- device_id
    display_block :- {[
    block :- B0 # Block B0 means whole screen
    images :- [
        {
            type :- Image/Video/GIF
            image:- Image_link,
            display_time :- -1
        }
    ]
    ]}
}
For Storing Multiple Image ( TO store playlist of multiple images or multiple static image to be shown in differnt grid block on the screen or playlist to be shown on multiple grid blocks on the screen)
This is for playlist to be shown on the whole screen

{
    id :- unique_Id
    type :- multiple
    device_id :- device_id
    display_block :- {[
    block :- B0 # Block B0 means whole screen
    images :- [
        {
            type :- Image/Video/GIF
            image:- Image_link,
            display_time :- 10 sec
        },
        {
            type :- Image/Video/GIF
            image:- Image_link,
            display_time :- 20 sec
        },
        ...
    ]
    ]}
}

This is for if the screen is sub divided into multiple blocks and each block will have a different image or playlist of images ( split_screen -> whole screen split into two halfs)
This means B1 will have a playlist and B2 will have a static image
{
    id :- unique_Id
    type :- multiple
    device_id :- device_id
    display_block :- {
    [
        block :- B1,
        images :- [
            {
                type :- Image/Video/GIF
                image:- Image_link,
                display_time :- 10 sec
            },
            {
                type :- Image/Video/GIF
                image:- Image_link,
                display_time :- 20 sec
            },
            ...
        ]
    ],
    [
        block :- B2,
        images :- [
            {
                type :- Image/Video/GIF
                image:- Image_link,
                display_time :- -1
            }
        ]
    ]
    }
}