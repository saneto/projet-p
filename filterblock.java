import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import org.springframework.beans.factory.annotation.Autowired;

public class UserContextFilter implements Filter {

    @Autowired
    private AuthenticationService authenticationService;

    @Override
    public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
            throws IOException, ServletException {

        // Cast to HttpServletRequest and HttpServletResponse to manipulate HTTP details
        HttpServletRequest httpRequest = (HttpServletRequest) request;
        HttpServletResponse httpResponse = (HttpServletResponse) response;

        try {
            // Retrieve the authenticated user from the service
            User currentUser = authenticationService.getAuthenticatedUser();

            if (currentUser != null) {
                // Set the user in the UserContext
                UserContext.setCurrentUser(currentUser);

                // Check if the request path is related to /groups and block non-admins
                String requestURI = httpRequest.getRequestURI();
                if (requestURI.startsWith("/groups")) {
                    if (!currentUser.getRoles().contains("ROLE_ADMIN")) {
                        // User is not an admin, block the request
                        httpResponse.setStatus(HttpServletResponse.SC_FORBIDDEN); // HTTP 403
                        httpResponse.getWriter().write("Access Denied: You must be an admin to access this resource.");
                        return;  // Stop further processing of the request
                    }
                }
            } else {
                // If no user is authenticated, you may also block access
                httpResponse.setStatus(HttpServletResponse.SC_UNAUTHORIZED); // HTTP 401
                httpResponse.getWriter().write("Unauthorized: You must be authenticated to access this resource.");
                return;
            }

            // Proceed with the filter chain if no issues
            chain.doFilter(request, response);

        } finally {
            // Clear the UserContext after the request
            UserContext.clear();
        }
    }

    @Override
    public void init(FilterConfig filterConfig) throws ServletException {}

    @Override
    public void destroy() {}
}