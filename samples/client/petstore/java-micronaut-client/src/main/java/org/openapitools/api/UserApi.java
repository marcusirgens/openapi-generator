/*
 * OpenAPI Petstore
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package org.openapitools.api;

import io.micronaut.http.annotation.*;
import io.micronaut.core.annotation.*;
import io.micronaut.http.client.annotation.Client;
import org.openapitools.query.QueryParam;
import io.micronaut.core.convert.format.Format;
import reactor.core.publisher.Mono;
import org.openapitools.model.User;
import javax.annotation.Generated;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import javax.validation.Valid;
import javax.validation.constraints.*;

@Generated(value="org.openapitools.codegen.languages.JavaMicronautClientCodegen")
@Client("${base-path}")
public interface UserApi {

  /**
   * Create user
   * This can only be done by the logged in user.
   *
   * @param _body Created user object (required)
   */
  @Post(uri="/user")
  @Produces(value={"*/*"})
  @Consumes(value={"application/json"})
  Mono<Object> createUser(
        @Body @Valid @NotNull User _body
  );

  /**
   * Creates list of users with given input array
   *
   * @param _body List of user object (required)
   */
  @Post(uri="/user/createWithArray")
  @Produces(value={"*/*"})
  @Consumes(value={"application/json"})
  Mono<Object> createUsersWithArrayInput(
        @Body @NotNull List<User> _body
  );

  /**
   * Creates list of users with given input array
   *
   * @param _body List of user object (required)
   */
  @Post(uri="/user/createWithList")
  @Produces(value={"*/*"})
  @Consumes(value={"application/json"})
  Mono<Object> createUsersWithListInput(
        @Body @NotNull List<User> _body
  );

  /**
   * Delete user
   * This can only be done by the logged in user.
   *
   * @param username The name that needs to be deleted (required)
   */
  @Delete(uri="/user/{username}")
  @Consumes(value={"application/json"})
  Mono<Object> deleteUser(
        @PathVariable(name="username") @NotNull String username
  );

  /**
   * Get user by user name
   *
   * @param username The name that needs to be fetched. Use user1 for testing. (required)
   * @return User
   */
  @Get(uri="/user/{username}")
  @Consumes(value={"application/json"})
  Mono<User> getUserByName(
        @PathVariable(name="username") @NotNull String username
  );

  /**
   * Logs user into the system
   *
   * @param username The user name for login (required)
   * @param password The password for login in clear text (required)
   * @return String
   */
  @Get(uri="/user/login")
  @Consumes(value={"application/json"})
  Mono<String> loginUser(
        @QueryParam(name="username") @NotNull String username, 
        @QueryParam(name="password") @NotNull String password
  );

  /**
   * Logs out current logged in user session
   *
   */
  @Get(uri="/user/logout")
  @Consumes(value={"application/json"})
  Mono<Object> logoutUser();

  /**
   * Updated user
   * This can only be done by the logged in user.
   *
   * @param username name that need to be deleted (required)
   * @param _body Updated user object (required)
   */
  @Put(uri="/user/{username}")
  @Produces(value={"*/*"})
  @Consumes(value={"application/json"})
  Mono<Object> updateUser(
        @PathVariable(name="username") @NotNull String username, 
        @Body @Valid @NotNull User _body
  );
}
