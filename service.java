import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;

@Service
public class AuthenticationService {

    @Autowired
    private UserRepository userRepository;  // Assume this exists for User entity

    public String getAuthenticatedUsername() {
        Authentication authentication = SecurityContextHolder.getContext().getAuthentication();
        return (authentication != null && authentication.isAuthenticated()) ? authentication.getName() : null;
    }

    public User getAuthenticatedUser() {
        String username = getAuthenticatedUsername();
        return username != null ? userRepository.findByUsername(username) : null;
    }
}