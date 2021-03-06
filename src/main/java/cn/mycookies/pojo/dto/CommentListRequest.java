package cn.mycookies.pojo.dto;

import cn.mycookies.common.constants.CommentTargetType;
import cn.mycookies.common.base.PageInfo4Request;
import com.fasterxml.jackson.databind.PropertyNamingStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.Data;

import javax.validation.constraints.NotNull;

/**
 * 评论列表请求参数
 *
 * @author Jann Lee
 * @date 2019-07-04 21:50
 **/
@Data
@ApiModel
@JsonNaming(PropertyNamingStrategy.SnakeCaseStrategy.class)
public class CommentListRequest extends PageInfo4Request {

    /**
     * 父级评论id
     */
    @ApiModelProperty("父级评论id，当查询一个留言下面的回复，对话时使用")
    private Integer pid;

    /**
     * 被评论主体类型
     */
    @ApiModelProperty("被评论主体类型")
    @NotNull
    private CommentTargetType targetType;
}
