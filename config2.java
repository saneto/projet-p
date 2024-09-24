import org.springframework.boot.web.servlet.FilterRegistrationBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class FilterConfig {

    @Bean
    public FilterRegistrationBean<UserContextFilter> userContextFilter(AuthenticationService authenticationService) {
        FilterRegistrationBean<UserContextFilter> registrationBean = new FilterRegistrationBean<>();
        
        // Create the filter and inject the AuthenticationService into it
        UserContextFilter filter = new UserContextFilter();
        filter.authenticationService = authenticationService;  // Inject service manually here

        registrationBean.setFilter(filter);
        registrationBean.addUrlPatterns("/*"); // Apply the filter to all URLs
        registrationBean.setOrder(1); // Set the order of this filter, e.g., high priority

        return registrationBean;
    }
}