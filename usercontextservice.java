import javax.servlet.Filter;
import javax.servlet.FilterChain;
import javax.servlet.FilterConfig;
import javax.servlet.ServletException;
import javax.servlet.ServletRequest;
import javax.servlet.ServletResponse;
import java.io.IOException;
import org.springframework.beans.factory.annotation.Autowired;

public class UserContextFilter implements Filter {

    @Autowired
    private AuthenticationService authenticationService;

    @Override
    public void doFilter(ServletRequest request, ServletResponse response, FilterChain chain)
            throws IOException, ServletException {
        try {
            User currentUser = authenticationService.getAuthenticatedUser();
            if (currentUser != null) {
                UserContext.setCurrentUser(currentUser);
            }
            chain.doFilter(request, response);
        } finally {
            // Clear the context after the request is done
            UserContext.clear();
        }
    }

    @Override
    public void init(FilterConfig filterConfig) throws ServletException {}

    @Override
    public void destroy() {}
}