import {GetCurrentAuthorizedUserRequest, SchemasUserPopulatedResponse} from "src/apiAutogenerated";
import {authService} from "src/service/services";

/**
 * Provides methods to interact with the Comments collection
 */
export class AuthService {

  /**
   * Call method for logout
   */
  public static async logOut() {
    await authService.logoutCurrentAuthorizedUser({provider: "google"});
  }

  /**
   * Get current authorized user data
   */
  public static async getCurrentUser(params: GetCurrentAuthorizedUserRequest): Promise<SchemasUserPopulatedResponse> {
    const user = await authService.getCurrentAuthorizedUser(params);

    return user;
  }

}
