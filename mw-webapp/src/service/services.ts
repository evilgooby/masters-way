import {
  AuthApi,
  CommentApi,
  CompositeWayApi,
  Configuration,
  DayReportApi,
  FavoriteUserApi,
  FavoriteUserWayApi,
  FetchParams,
  FromUserMentoringRequestApi,
  JobDoneApi,
  JobDoneJobTagApi,
  JobTagApi,
  MentorUserWayApi,
  MetricApi,
  Middleware,
  PlanApi,
  PlanJobTagApi,
  ProblemApi,
  ProblemJobTagApi,
  ToUserMentoringRequestApi,
  UserApi, UserTagApi,
  WayApi,
  WayCollectionApi,
  WayCollectionWayApi,
  WayTagApi,
} from "src/apiAutogenerated";
import {userStore} from "src/globalStore/UserStore";
import {isAuthCookieExists} from "src/utils/cookieUtils";
import {env} from "src/utils/env/env";

const checkAuthMiddleware: Middleware = {

  /**
   * Pre request middle ware
   * @param context
   */
  pre: (): Promise<FetchParams | void> => {
    const isSessionValid = isAuthCookieExists();
    if (!isSessionValid) {
      userStore.clearUser();
    }

    return Promise.resolve(undefined);
  },

  /**
   * Pre request middle ware
   * @param context
   */
  post: (): Promise<Response | void> => {
    const isSessionValid = isAuthCookieExists();
    if (!isSessionValid) {
      userStore.clearUser();
    }

    return Promise.resolve(undefined);
  },
};

const configuration = new Configuration({
  basePath: env.API_BASE_PATH,
  credentials: "include",
  middleware: [checkAuthMiddleware],
});

export const authService = new AuthApi(configuration);
export const userService = new UserApi(configuration);
export const wayCollectionService = new WayCollectionApi(configuration);
export const favoriteUserService = new FavoriteUserApi(configuration);
export const userTagService = new UserTagApi(configuration);
export const wayTagService = new WayTagApi(configuration);
export const wayCollectionWayService = new WayCollectionWayApi(configuration);
export const wayService = new WayApi(configuration);
export const toUserMentoringRequestService = new ToUserMentoringRequestApi(configuration);
export const problemJobTagService = new ProblemJobTagApi(configuration);
export const planJobTagService = new PlanJobTagApi(configuration);
export const jobDoneJobTagService = new JobDoneJobTagApi(configuration);
export const jobTagService = new JobTagApi(configuration);
export const jobDoneService = new JobDoneApi(configuration);
export const planService = new PlanApi(configuration);
export const problemService = new ProblemApi(configuration);
export const commentService = new CommentApi(configuration);
export const dayReportService = new DayReportApi(configuration);
export const metricService = new MetricApi(configuration);
export const favoriteUserWayService = new FavoriteUserWayApi(configuration);
export const fromUserMentoringRequest = new FromUserMentoringRequestApi(configuration);
export const mentorUserWay = new MentorUserWayApi(configuration);
export const compositeWayService = new CompositeWayApi(configuration);
