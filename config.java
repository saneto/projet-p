import org.springframework.boot.web.servlet.FilterRegistrationBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class FilterConfig {

    @Bean
    public FilterRegistrationBean<UserContextFilter> userContextFilter() {
        FilterRegistrationBean<UserContextFilter> registrationBean = new FilterRegistrationBean<>();
        
        // Instantiate and register your filter
        registrationBean.setFilter(new UserContextFilter());
        
        // Set URL patterns the filter should be applied to (optional)
        registrationBean.addUrlPatterns("/*"); // Apply to all URLs
        
        // Optionally set the filter order
        registrationBean.setOrder(1); // Set the order of this filter

        return registrationBean;
    }
}