-- name: FeedLikeExists :one
SELECT EXISTS (
    SELECT 1 FROM likes
    WHERE user_id = @user_id
    AND feed_id = @feed_id
    AND entity_type = 'feed'
);

-- name: CommentLikeExists :one
SELECT EXISTS (
    SELECT 1 FROM likes
    WHERE user_id = @user_id
    AND comment_id = @comment_id
    AND entity_type = 'comment'
);


-- name: CreateFeedLike :exec
INSERT INTO likes  
(user_id, feed_id, entity_type)
VALUES (@user_id, @feed_id, 'feed');

-- name: CreateCommentLike :exec
INSERT INTO likes  
(user_id, comment_id, entity_type)
VALUES (@user_id, @comment_id, 'comment');


-- name: DeleteFeedLike :exec
DELETE FROM likes
WHERE user_id = @user_id
AND feed_id = @feed_id
AND entity_type = 'feed';

-- name: DeleteCommentLike :exec
DELETE FROM likes
WHERE user_id = @user_id
AND comment_id = @comment_id
AND entity_type = 'comment';