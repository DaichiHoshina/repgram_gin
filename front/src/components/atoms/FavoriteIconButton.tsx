import { CardActions, IconButton, Typography } from "@material-ui/core";
import React, { FC, useEffect, useState } from "react";
import FavoriteIcon from "@material-ui/icons/Favorite";
import { TPresentation } from "../../modules/Presentation";
import { useDispatch } from "react-redux";
import { TUser } from "../../modules/User";
import { createLike, deleteLike } from "../../services/Like";
import { fetchPresentations } from "../../services/Presentation";

interface Props {
  presentation?: TPresentation;
  loginUser?: TUser;
  key?: number;
  page?: number;
}

const FavoriteIconButton: FC<Props> = ({
  presentation,
  loginUser,
  page,
}: Props) => {
  const dispatch = useDispatch();

  const [isLike, setIsLike] = useState(false);

  const handleClick = () => {
    console.log("-------------------------------------------------");
    console.log("いいねボタン押下");
    console.log("-------------------------------------------------");
    setIsLike(!isLike);
    if (isLike) {
      likeDelete(loginUser?.id!, presentation?.id!);
    } else {
      likeCreate(loginUser?.id!, presentation?.id!);
    }
  };

  useEffect(() => {
    presentation?.likes?.some((like) => {
      if (like.user_id! == loginUser?.id!) {
        setIsLike(true);
      }
    });
  }, []);

  const likeCreate = async (
    user_id: number | string,
    presentation_id: number | string
  ) => {
    const likeValue = {
      user_id: user_id,
      presentation_id: presentation_id,
    };
    await dispatch(
      createLike({
        like: likeValue,
      })
    );
    await dispatch(
      fetchPresentations({
        page: page,
        per: 6,
      })
    );
  };

  const likeDelete = async (
    user_id: number | string,
    presentation_id: number | string
  ) => {
    const likeValue = {
      user_id: user_id,
      presentation_id: presentation_id,
    };
    await dispatch(
      deleteLike({
        like: likeValue,
      })
    );
    await dispatch(
      fetchPresentations({
        page: page,
        per: 6,
      })
    );
  };

  return (
    <>
      <CardActions disableSpacing>
        <IconButton
          aria-label="add to favorites"
          disabled={!loginUser}
          onClick={handleClick}
        >
          <FavoriteIcon color={isLike ? "secondary" : "disabled"} />
          <div className="ml-1">
            <Typography color={isLike ? "" : "error"}>
              {presentation?.likes?.length}
            </Typography>
          </div>
        </IconButton>
      </CardActions>
    </>
  );
};
export default FavoriteIconButton;
